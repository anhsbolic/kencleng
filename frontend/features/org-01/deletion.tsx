'use client';
import Link from 'next/link';
import { useEffect, useRef, useState } from 'react';
import { clearMarker } from './attempt';
import { request, message, types, type Schema } from './client';
import { MockControls, useGoogleNavigation } from './provider';
export function DeletionReview() {
    const [session, setSession] = useState<Schema<'SessionResponse'>['data'] | null>(null), [preview, setPreview] = useState<Schema<'DeletionPreviewResponse'>['data'] | null>(null), [notice, setNotice] = useState(''), [busy, setBusy] = useState(false), [terminal, setTerminal] = useState(false), [confirmed, setConfirmed] = useState(false), [proofNeeded, setProofNeeded] = useState(false), [result, setResult] = useState<Schema<'DeletionResponse'>['data'] | null>(null);
    const gate = useRef(false), status = useRef<HTMLDivElement>(null), navigate = useGoogleNavigation();
    function announce(text: string) { setNotice(text); requestAnimationFrame(() => status.current?.focus()); }
    async function getPreview() { if (gate.current || terminal)
        return; gate.current = true; setBusy(true); setPreview(null); setConfirmed(false); try {
        const current = await request<Schema<'SessionResponse'>>('getSession');
        if (current.kind !== 'ok') {
            setSession(null);
            announce(message(current, 'Pemeriksaan akun'));
            return;
        }
        setSession(current.data.data);
        const r = await request<Schema<'DeletionPreviewResponse'>>('createAccountDeletionPreview', { body: {}, csrf: current.data.data.csrfToken });
        if (r.kind === 'ok') {
            setPreview(r.data.data);
            setProofNeeded(false);
            announce('Pratinjau siap ditinjau. Belum ada penghapusan yang diterima.');
        }
        else {
            setProofNeeded(r.kind === 'problem' && r.problem.type === types.proof);
            announce(message(r, 'Pratinjau penghapusan'));
        }
    }
    finally {
        gate.current = false;
        setBusy(false);
    } }
    useEffect(() => { void request<Schema<'SessionResponse'>>('getSession').then(r => { if (r.kind === 'ok')
        setSession(r.data.data);
    else
        announce(message(r, 'Pemeriksaan sesi')); }); }, []);
    async function prove() { if (!session || gate.current || terminal)
        return; gate.current = true; setBusy(true); setPreview(null); setConfirmed(false); try {
        const r = await request<Schema<'AttemptResponse'>>('createGoogleAttempt', { body: { purpose: 'deleteAccount' }, csrf: session.csrfToken });
        if (r.kind === 'ok') {
            try {
                await navigate(r.data.data.authorizationUrl, true);
            }
            catch {
                announce('Tujuan verifikasi tidak dapat dipastikan. Tidak ada pengalihan.');
            }
        }
        else
            announce(message(r, 'Verifikasi akun'));
    }
    finally {
        gate.current = false;
        setBusy(false);
    } }
    async function erase() {
        if (!session || !preview || !confirmed || gate.current || terminal)
            return;
        if (Date.now() >= Date.parse(preview.expiresAt)) {
            setPreview(null);
            announce('Pratinjau kedaluwarsa. Tinjau ulang sebelum mengonfirmasi.');
            return;
        }
        gate.current = true;
        setBusy(true);
        try {
            const current = await request<Schema<'SessionResponse'>>('getSession');
            if (current.kind !== 'ok' || current.data.data.accountId.toLowerCase() !== session.accountId.toLowerCase()) {
                setPreview(null);
                setSession(null);
                announce('Akun belum cocok. Tidak ada konfirmasi yang dikirim. Periksa sesi dan cakupan kembali.');
                return;
            }
            const r = await request<Schema<'DeletionResponse'>>('deleteAccount', { body: { previewId: preview.previewId, closeOrganizations: true }, csrf: current.data.data.csrfToken });
            setPreview(null);
            setConfirmed(false);
            if (r.kind === 'ok') {
                clearMarker();
                setSession(null);
                setTerminal(true);
                setResult(r.data.data);
                announce('Akses berakhir. Data pada sistem aktif telah dihapus sesuai respons sistem.');
            }
            else if (r.kind === 'unknown' || (r.kind === 'problem' && (r.problem.status >= 500 || r.problem.status === 401))) {
                clearMarker();
                setSession(null);
                setTerminal(true);
                announce(message(r, 'Konfirmasi penghapusan') + ' Jangan masuk untuk memeriksa akun lama atau mengirim konfirmasi baru.');
            }
            else {
                setProofNeeded(r.problem.type === types.proof);
                announce(message(r, 'Konfirmasi penghapusan'));
            }
        }
        finally {
            gate.current = false;
            setBusy(false);
        }
    }
    return <><header className="site-header"><Link href="/" className="brand">Kencleng</Link><nav><Link href="/organizations">Organisasi</Link><Link href="/privacy">Privasi</Link></nav></header><main className="narrow"><div className="eyebrow">Kendali akun</div><h1>Tinjau penghapusan akun</h1><div ref={status} role="status" aria-live="polite" tabIndex={-1} className={notice ? 'notice' : 'status-empty'}>{notice}</div>{!terminal && <><p className="lead">Pahami seluruh konteks yang ikut ditutup sebelum mengambil keputusan.</p><p>Penutupan di Kencleng bukan pembubaran organisasi secara hukum. Tidak ada pemindahan Owner atau penerus otomatis. Jika organisasi perlu tetap berjalan atau memiliki relasi/aktivitas tambahan, perlukan penilaian pengelola privasi.</p><p>Akun yang bertindak: {session ? <code>{session.accountId}</code> : 'belum dapat diperiksa'}</p><p>Verifikasi melalui Google harus memakai akun yang sama. Ini tidak menjanjikan Google meminta password atau MFA. Verifikasi saja tidak menghapus akun atau membuat organisasi.</p><div className="actions"><button disabled={busy || !session} onClick={getPreview}>Periksa cakupan penghapusan</button><button className="secondary" disabled={busy || !session} onClick={prove}>{proofNeeded ? 'Verifikasi akun yang sama dengan Google' : 'Verifikasi ulang melalui Google'}</button></div></>}
 {preview && !terminal && <section className="deletion-scope"><h2>Seluruh organisasi yang akan ditutup</h2>{preview.organizations.length === 0 ? <p>Akun ini tidak memiliki organisasi dalam cakupan pratinjau.</p> : <ul>{preview.organizations.map(org => <li key={org.id}><strong>{org.name}</strong><p className="small">Nama dari pengelola · Owner internal</p><code>{org.id}</code></li>)}</ul>}<p>Pratinjau berlaku sampai {new Date(preview.expiresAt).toLocaleString('id-ID')}. Perubahan akun atau cakupan memerlukan pemeriksaan baru.</p><h3>Konsekuensi sebelum konfirmasi</h3><p>Konfirmasi menutup seluruh konteks di atas, mengakhiri akses semua sesi, dan menghapus data akun serta relasi pada sistem aktif. Akses dapat ditahan saat penghapusan diproses. Pembatalan otomatis tidak tersedia setelah niat diterima.</p><p>Salinan cadangan belum langsung hilang. Batas umur maksimum 14 hari adalah kewajiban retensi, bukan bukti penghapusan sudah dijalankan. Penyelesaian seluruh salinan memerlukan bukti tersendiri dan pemulihan tidak boleh menghidupkan kembali akses/data yang dihapus.</p><label className="confirmation"><input type="checkbox" checked={confirmed} onChange={e => setConfirmed(e.target.checked)} disabled={busy}/>Saya memilih menutup seluruh konteks yang tercantum dan menghapus akun ini, termasuk bila daftar organisasi kosong.</label><div className="actions"><button className="danger" disabled={busy || !confirmed} onClick={erase}>Tutup konteks dan hapus akun</button><button className="secondary" disabled={busy} onClick={() => { setPreview(null); setConfirmed(false); announce('Tidak ada konfirmasi penghapusan yang dikirim.'); }}>Batalkan sebelum konfirmasi</button></div></section>}
 {result && <section className="result"><h2>Fakta yang telah dikonfirmasi</h2><ul><li>Akses akun berakhir.</li><li>Data sistem aktif dihapus.</li><li>{result.remainingCopies === 'pendingVerification' ? 'Salinan lain: penyelesaian belum terverifikasi.' : 'Tidak ada salinan tersisa menurut bukti disposal sistem.'}</li></ul></section>}
 <section className="follow-up"><h2>{terminal ? 'Tindak lanjut hasil' : 'Organisasi perlu tetap berjalan?'}</h2><p>Pengelola privasi Yayasan Kencleng Indonesia bertanggung jawab atas penilaian dan tindak lanjut. Kanal yang ditetapkan adalah care@kencleng.dev. Operasional inbox, verifikasi kontrol akun, tindak lanjut, dan bukti penyelesaian belum diverifikasi untuk production.</p><p>Jangan kirim data pengguna nyata dalam pengembangan ini. Email terkirim bukan bukti permintaan diterima; tidak ada janji email otomatis atau tenggat yang belum disetujui.</p><Link href="/privacy">Baca pemberitahuan lengkap — DRAFT</Link></section><MockControls /></main></>;
}
