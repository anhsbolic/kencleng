'use client';
import Link from 'next/link';
import { useCallback, useEffect, useRef, useState } from 'react';
import { useForm } from 'react-hook-form';
import { request, message, types, type Schema, type Result } from './client';
import { canonicalName, validName, requestId, readMarker, saveMarker, clearMarker, refreshSession, type Attempt } from './attempt';
import { MockControls, useGoogleNavigation } from './provider';
export const ownerCopy = 'Anda menjadi Owner awal di Kencleng. Ini tidak memverifikasi nama, mewakili organisasi secara hukum, atau memberi peran Reviewer.';
export function Workspace({ continuation = false }: {
    continuation?: boolean;
}) {
    const [session, setSession] = useState<Schema<'SessionResponse'>['data'] | null>(null), [checked, setChecked] = useState(false), [busy, setBusy] = useState(false), [notice, setNotice] = useState(''), [review, setReview] = useState<string | null>(null), [attempt, setAttempt] = useState<Attempt | null>(null), [success, setSuccess] = useState<Schema<'CreationResponse'>['data'] | null>(null), [list, setList] = useState<Schema<'OrganizationList'> | null>(null), [detail, setDetail] = useState<Schema<'OrganizationResponse'>['data'] | null>(null), [listNotice, setListNotice] = useState(''), [loginWarning, setLoginWarning] = useState(false);
    const gate = useRef(false), status = useRef<HTMLDivElement>(null), epoch = useRef(0), principal = useRef<string | null>(null), navigate = useGoogleNavigation();
    const { register, handleSubmit, getValues, formState: { errors }, reset, setError } = useForm<{
        name: string;
    }>({ defaultValues: { name: '' } });
    const announce = (text: string) => { setNotice(text); requestAnimationFrame(() => status.current?.focus()); };
    const loadList = useCallback(async (cursor?: string) => { const generation = epoch.current, account = principal.current; if (!account)
        return; const r = await request<Schema<'OrganizationList'>>('listOrganizations', { cursor }); if (generation !== epoch.current || account !== principal.current)
        return; if (r.kind === 'ok' && r.data.data.some(p => p.owner.accountId.toLowerCase() !== account?.toLowerCase())) {
        setList(null);
        setListNotice('Daftar akun tidak konsisten; tidak ditampilkan.');
        return;
    } if (r.kind === 'ok') {
        setList(old => cursor && old ? { ...r.data, data: [...old.data, ...r.data.data] } : r.data);
        setListNotice('');
    }
    else {
        if (!cursor)
            setList(null);
        setListNotice(message(r, 'Pemeriksaan organisasi'));
    } }, []);
    const applySession = useCallback((r: Result<Schema<'SessionResponse'>>) => {
        setChecked(true);
        if (r.kind === 'ok') {
            const current = r.data.data;
            const marker = readMarker();
            if ((principal.current && principal.current.toLowerCase() !== current.accountId.toLowerCase()) || (marker && marker.accountId.toLowerCase() !== current.accountId.toLowerCase())) {
                clearMarker();
                setAttempt(null);
                setSuccess(null);
                setDetail(null);
                setList(null);
                setReview(null);
                reset();
                setNotice('Akun berubah. Data pribadi dan referensi permintaan sebelumnya dibersihkan dari tab ini.');
            }
            else if (marker)
                setAttempt(old => old ?? marker);
            if (principal.current !== current.accountId)
                epoch.current++;
            principal.current = current.accountId;
            setSession(current);
            void loadList();
        }
        else {
            epoch.current++;
            principal.current = null;
            setSession(null);
            setList(null);
            setDetail(null);
            if (r.kind !== 'problem' || r.problem.type !== types.unauthenticated)
                setNotice(message(r, 'Pemeriksaan sesi'));
        }
        return r;
    }, [loadList, reset]);
    useEffect(() => { void request<Schema<'SessionResponse'>>('getSession').then(applySession); }, [applySession]);
    async function action(run: () => Promise<void>) { if (gate.current)
        return; gate.current = true; setBusy(true); try {
        await run();
    }
    finally {
        gate.current = false;
        setBusy(false);
    } }
    async function login(proceed = false) { if (getValues('name') && !proceed) {
        setLoginWarning(true);
        return;
    } setLoginWarning(false); await action(async () => { if (getValues('name'))
        announce('Nama di formulir tidak disimpan saat berpindah ke Google. Permintaan yang belum pasti tetap perlu diperiksa.'); const r = await request<Schema<'AttemptResponse'>>('createGoogleAttempt', { body: {} }); if (r.kind === 'ok') {
        try {
            await navigate(r.data.data.authorizationUrl, false);
        }
        catch {
            announce('Tujuan masuk tidak dapat diverifikasi. Tidak ada pengalihan.');
        }
    }
    else
        announce(message(r, 'Memulai masuk')); }); }
    async function create(a: Attempt) {
        if (!session || !a.name)
            return;
        await action(async () => {
            const fresh = await request<Schema<'SessionResponse'>>('getSession');
            if (fresh.kind !== 'ok' || fresh.data.data.accountId.toLowerCase() !== a.accountId.toLowerCase()) {
                applySession(fresh);
                announce('Akun belum cocok. Permintaan tidak dikirim. Periksa sesi dan hasil sebelumnya.');
                return;
            }
            setAttempt(a);
            setReview(null);
            const durable = saveMarker(a);
            const r = await request<Schema<'CreationResponse'>>('createOrganization', { body: { name: a.name }, key: a.requestId, csrf: fresh.data.data.csrfToken });
            if (r.kind === 'ok' && r.data.data.requestId.toLowerCase() === a.requestId.toLowerCase() && r.data.data.owner.accountId.toLowerCase() === a.accountId.toLowerCase() && r.data.data.organization.name === a.name) {
                setSuccess(r.data.data);
                setAttempt(null);
                clearMarker();
                reset();
                announce('Organisasi dan Owner awal tercatat. Nama merupakan informasi dari pengelola.');
                await loadList();
            }
            else {
                if (r.kind === 'problem' && r.problem.type === types.validation) {
                    setAttempt(null);
                    clearMarker();
                    setError('name', { type: 'server', message: 'Periksa nama: 1–200 karakter, tanpa karakter kontrol.' });
                }
                announce((r.kind === 'ok' ? 'Respons tidak konsisten; hasil belum dapat dipastikan.' : message(r, 'Pembuatan organisasi')) + (!durable ? ' Penyimpanan tab tidak tersedia; referensi permintaan akan hilang saat muat ulang.' : ''));
            }
        });
    }
    async function checkResult() { if (!session || !attempt)
        return; await action(async () => { const fresh = await request<Schema<'SessionResponse'>>('getSession'); if (fresh.kind !== 'ok' || fresh.data.data.accountId.toLowerCase() !== attempt.accountId.toLowerCase()) {
        applySession(fresh);
        announce('Akun harus cocok sebelum hasil diperiksa.');
        return;
    } const r = await request<Schema<'CreationResponse'>>('getOrganizationCreationResult', { id: attempt.requestId }); if (r.kind === 'ok' && r.data.data.requestId === attempt.requestId && r.data.data.owner.accountId.toLowerCase() === attempt.accountId.toLowerCase()) {
        setSuccess(r.data.data);
        setAttempt(null);
        clearMarker();
        announce('Hasil permintaan ditemukan: organisasi dan Owner awal tercatat.');
        await loadList();
    }
    else
        announce(r.kind === 'ok' ? 'Hasil tidak konsisten; belum dapat dipastikan.' : message(r, 'Pemeriksaan hasil')); }); }
    async function recover() { await action(async () => { const current = await request<Schema<'SessionResponse'>>('getSession'); if (current.kind !== 'ok') {
        applySession(current);
        announce(message(current, 'Pemulihan sesi'));
        return;
    } const old = principal.current; const refreshed = await refreshSession(current.data.data.csrfToken); applySession(refreshed as Result<Schema<'SessionResponse'>>); announce(refreshed.kind === 'ok' && (!old || refreshed.data?.data.accountId === old) ? 'Sesi diperiksa kembali. Tidak ada permintaan organisasi atau penghapusan yang dikirim ulang.' : 'Pemulihan sesi belum pasti. Jangan mengulang rotasi secara otomatis; pilih masuk kembali secara eksplisit.'); }); }
    async function logout() { await action(async () => { epoch.current++; principal.current = null; const r = await request<void>('logoutSession', { body: {}, csrf: session?.csrfToken }); setSession(null); setList(null); setDetail(null); setSuccess(null); setAttempt(null); setReview(null); principal.current = null; clearMarker(); reset(); announce(r.kind === 'ok' ? 'Akses sesi berakhir. Organisasi yang sudah tercatat tidak dibatalkan.' : 'Data pribadi di tab dibersihkan. Pencabutan sesi di server belum terbukti; hasil permintaan sebelumnya tetap tidak pasti.'); }); }
    return <>
 <header className="site-header"><Link href="/" className="brand">Kencleng</Link><nav aria-label="Navigasi utama"><Link href="/organizations">Organisasi</Link><Link href="/privacy">Privasi</Link>{session && <button className="text-button" disabled={busy} onClick={logout}>Keluar</button>}</nav></header>
 <main className="workspace"><div className="eyebrow">Konteks organisasi</div><h1>{session ? 'Organisasi Anda' : 'Mulai dari satu organisasi.'}</h1><div ref={status} tabIndex={-1} role="status" aria-live="polite" className={notice ? 'notice' : 'status-empty'}>{notice}</div>
 {!checked ? <p role="status">Memeriksa sesi…</p> : !session ? <div className="start-grid"><section><p className="lead">Satu ruang untuk organisasi, dengan tanggung jawab yang jelas.</p><ol className="steps"><li><strong>Masuk dengan akun Anda</strong><p>Akun Kencleng menjadi dasar atribusi. Masuk tidak membuat organisasi.</p></li><li><strong>Tinjau nama dan peran awal</strong><p>{ownerCopy}</p></li><li><strong>Periksa hasilnya kembali</strong><p>Organisasi tetap dapat diperiksa melalui identitasnya, termasuk bila namanya serupa.</p></li></ol></section><section className="signin-panel"><h2>Sebelum melanjutkan</h2><p>Yayasan Kencleng Indonesia mengelola identitas akun, nama organisasi, dan relasi Owner untuk konteks internal Kencleng.</p><p>Google digunakan untuk masuk dengan cakupan openid. Tombol masuk bukan persetujuan menyeluruh atas semua pemrosesan.</p><p><Link href="/privacy">Baca pemberitahuan lengkap — DRAFT</Link>. Dasar hukum, penyedia/lokasi lain, serta kesiapan penanganan privasi masih ditahan sebelum production. Kanal yang ditetapkan: care@kencleng.dev; operasionalnya belum terverifikasi.</p><button disabled={busy} onClick={() => login()}>Lanjutkan dengan Google</button><p className="small">Pengembangan memakai data sintetis saja.</p>{loginWarning && <div className="notice"><p>Nama di formulir akan hilang saat berpindah ke Google. Referensi permintaan yang belum pasti tetap perlu diperiksa.</p><button onClick={() => login(true)}>Saya memahami, lanjutkan masuk</button><button className="secondary" onClick={() => setLoginWarning(false)}>Kembali</button></div>}{continuation && <p>Kelanjutan masuk tidak mengirim permintaan organisasi.</p>}</section></div> : <>
 <div className="account-line"><span>Akun yang bertindak</span><code>{session.accountId}</code><button className="text-button" disabled={busy} onClick={recover}>Periksa dan pulihkan sesi</button></div>
 {attempt && <section className="recovery" aria-labelledby="recovery-title"><h2 id="recovery-title">Hasil permintaan perlu diperiksa</h2><p>Jangan menganggap efek dibatalkan. Periksa hasil permintaan ini atau organisasi Anda sebelum memulai permintaan lain.</p><p className="small">Referensi permintaan <code>{attempt.requestId}</code></p><div className="actions"><button disabled={busy} onClick={checkResult}>Periksa hasil permintaan</button>{attempt.name && <button className="secondary" disabled={busy} onClick={() => create(attempt)}>Kirim ulang nama dan kunci yang sama</button>}</div>{!attempt.name && <p>Nama tidak disimpan saat muat ulang. Pengiriman ulang persis tidak tersedia; jangan menebak nama berdasarkan daftar.</p>}<details><summary>Lepaskan referensi dari tab ini</summary><p>Ini tidak membatalkan efek di server. Permintaan baru dapat membuat organisasi lain, termasuk dengan nama yang sama.</p><button className="secondary" disabled={busy} onClick={() => { clearMarker(); setAttempt(null); announce('Referensi dilepas dari tab. Efek permintaan sebelumnya tidak dibatalkan.'); }}>Saya memahami, lepaskan referensi</button></details></section>}
 {success && <section className="result"><h2>Organisasi tercatat</h2><h3>{success.organization.name}</h3><p>Nama dari pengelola · Owner internal: akun Anda.</p><code>{success.organization.id}</code><p>Ini tidak menyatakan verifikasi hukum, persetujuan Reviewer, publikasi, atau penerimaan donasi.</p></section>}
 <div className="product-grid"><section><h2>Buat organisasi</h2><p>Nama tidak harus unik. Organisasi dengan nama serupa tetap memiliki identitas terpisah.</p>{review !== null ? <div><h3 id="review-heading" tabIndex={-1}>Tinjau sebelum membuat</h3><dl><dt>Nama dari pengelola</dt><dd>{review}</dd><dt>Akun yang menjadi Owner awal</dt><dd><code>{session.accountId}</code></dd></dl><p>{ownerCopy}</p><div className="actions"><button disabled={busy || !!attempt} onClick={() => create({ accountId: session.accountId, requestId: requestId(), name: review })}>Buat organisasi dengan nama ini</button><button disabled={busy} className="secondary" onClick={() => { setReview(null); requestAnimationFrame(() => document.getElementById('org-name')?.focus()); }}>Kembali mengubah nama</button></div></div> : <form onSubmit={handleSubmit(({ name }) => { setReview(canonicalName(name)); setSuccess(null); requestAnimationFrame(() => document.getElementById('review-heading')?.focus()); })}><label htmlFor="org-name">Nama organisasi</label><input id="org-name" {...register('name', { validate: n => validName(n) || 'Masukkan 1–200 karakter, tanpa karakter kontrol.' })} aria-invalid={!!errors.name} aria-describedby="name-help name-error" disabled={busy || !!attempt}/><p id="name-help" className="small">Gunakan nama yang ingin ditampilkan. Huruf besar dan spasi di tengah dipertahankan.</p><p id="name-error" role="alert" className="field-error">{errors.name?.message}</p><button disabled={busy || !!attempt} type="submit">Tinjau nama dan Owner awal</button></form>}</section>
 <section className="organization-list"><div className="section-heading"><h2>Daftar saat ini</h2><button className="text-button" disabled={busy} onClick={() => loadList()}>Muat ulang daftar</button></div>{listNotice ? <p role="alert">{listNotice}</p> : list?.data.length === 0 ? <p>Belum ada organisasi yang terlihat untuk akun ini.</p> : !list ? <p>Memuat organisasi…</p> : <><ul>{list.data.map(pair => <li key={pair.organization.id}><h3>{pair.organization.name}</h3><p className="small">Nama dari pengelola · Owner internal</p><code>{pair.organization.id}</code><button className="text-button" disabled={busy} onClick={async () => { const generation = epoch.current, account = principal.current; if (!account)
            return; const r = await request<Schema<'OrganizationResponse'>>('getOrganization', { id: pair.organization.id }); if (generation !== epoch.current || account !== principal.current)
            return; if (r.kind === 'ok' && r.data.data.owner.accountId.toLowerCase() === account.toLowerCase()) {
            setDetail(r.data.data);
            announce('Detail dan relasi saat ini tersedia di bawah daftar.');
        }
        else {
            setDetail(null);
            announce(message(r, 'Detail organisasi'));
        } }}>Periksa detail saat ini</button></li>)}</ul>{list.meta && <button className="secondary" disabled={busy} onClick={() => loadList(list.meta?.nextCursor)}>Muat organisasi berikutnya</button>}</>}{detail && <div className="detail"><h3>Relasi saat ini</h3><p>{detail.organization.name}</p><code>{detail.organization.id}</code><p>Akun <code>{detail.owner.accountId}</code> · {detail.owner.role} · internal</p><p>Dibuat {new Date(detail.organization.createdAt).toLocaleString('id-ID')}</p></div>}</section></div>
 <section className="privacy-link"><h2>Kendali atas akun Anda</h2><p>Penutupan konteks dan penghapusan akun memerlukan pemeriksaan cakupan serta konfirmasi eksplisit.</p><Link href="/account-deletion/review">Tinjau penghapusan akun</Link></section></>}
 <MockControls /></main><footer>Kencleng · konteks internal organisasi <Link href="/privacy">Pemberitahuan pemrosesan data</Link></footer></>;
}
