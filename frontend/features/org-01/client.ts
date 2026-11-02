import type { ZodType } from 'zod';
import { operations } from './generated/validation';
import type { components } from './generated/api';
export type Schema<K extends keyof components['schemas']> = components['schemas'][K];
export type Problem = components['schemas'][Extract<keyof components['schemas'], `Problem${string}`>];
export type Result<T> = {
    kind: 'ok';
    data: T;
} | {
    kind: 'problem';
    problem: Problem;
    wait?: number;
} | {
    kind: 'unknown';
};
export const types = {
    pending: 'urn:uuid:6d143e2a-e3ba-4e71-8c67-1c8f7d4bae85', proof: 'urn:uuid:0ddfd153-87b2-4619-a11e-f48418401c21', unauthenticated: 'urn:uuid:0c56be0a-40a6-45c9-bb49-c20b075ba28b', validation: 'urn:uuid:1b1d2eb4-2c68-4809-8910-f40bd94e1e42', notObserved: 'urn:uuid:cbf15221-1fdb-47ca-8b8a-8285b9e031b1', expired: 'urn:uuid:22166aa9-86c0-4033-a5c6-49f7cb83f9ff', scope: 'urn:uuid:bd8b4ade-aaef-4ea9-9d73-2585c7b31ea6', preview: 'urn:uuid:25c9699d-d474-496b-8214-65ceb830f197', manual: 'urn:uuid:5a799183-67f6-49a9-90cd-ae9e06af1cf2', rate: 'urn:uuid:364a031e-17d4-42d0-af70-06c79716e98e', mismatch: 'urn:uuid:144ea3b4-83cc-4b0f-961a-892f3fc09b86', forbidden: 'urn:uuid:142cb767-7ce1-403f-88b4-b1fe7c6c2f82', invalid: 'urn:uuid:9ed32605-5d10-4f59-b811-042f45fdffed', unavailable: 'urn:uuid:bc32fcd9-a6ce-4c42-bfa2-d09431017922', busy: 'urn:uuid:cc31e0af-5ded-4ccd-afe6-cb0c592323cc', resultUnavailable: 'urn:uuid:59625f7b-1a32-46af-aae7-c42ddf7be1fd', resultUnknown: 'urn:uuid:f67a74ed-f3a2-45bb-a2ee-8b14e9c4c96c', resource: 'urn:uuid:5ed0b852-a063-4c56-85a7-a0d7c1775000'
} as const;
export async function request<T>(operation: keyof typeof operations, options: {
    body?: unknown;
    csrf?: string;
    id?: string;
    cursor?: string;
    key?: string;
} = {}): Promise<Result<T>> {
    const op = operations[operation];
    let path: string = op.path;
    if (options.id)
        path = path.replace(/\{[^}]+\}/, encodeURIComponent(options.id));
    if (options.cursor)
        path += '?cursor=' + encodeURIComponent(options.cursor);
    try {
        const response = await fetch('/api' + path, { method: op.method, credentials: 'same-origin', cache: 'no-store', redirect: 'error', headers: { ...(options.body === undefined ? {} : { 'Content-Type': 'application/json' }), ...(options.csrf ? { 'X-CSRF-Token': options.csrf } : {}), ...(options.key ? { 'Idempotency-Key': options.key } : {}) }, ...(options.body === undefined ? {} : { body: JSON.stringify(options.body) }), signal: AbortSignal.timeout(15000) });
        const validator = (op.responses as Record<number, ZodType>)[response.status];
        if (!validator)
            return { kind: 'unknown' };
        if (response.status === 204)
            return { kind: 'ok', data: undefined as T };
        const media = response.headers.get('Content-Type')?.split(';')[0];
        if (media !== (response.ok ? 'application/json' : 'application/problem+json'))
            return { kind: 'unknown' };
        const text = await response.text();
        if (new TextEncoder().encode(text).length > 65536)
            return { kind: 'unknown' };
        const parsed = validator.safeParse(JSON.parse(text));
        if (!parsed.success)
            return { kind: 'unknown' };
        if (response.ok)
            return { kind: 'ok', data: parsed.data as T };
        const wait = Number(response.headers.get('Retry-After'));
        return { kind: 'problem', problem: parsed.data as Problem, ...(Number.isInteger(wait) && wait >= 1 && wait <= 60 ? { wait } : {}) };
    }
    catch {
        return { kind: 'unknown' };
    }
}
export function message(result: Result<unknown>, context = 'Permintaan'): string {
    if (result.kind === 'ok')
        return '';
    if (result.kind === 'unknown')
        return `${context}: hasil belum dapat dipastikan. Jangan menganggap efek dibatalkan.`;
    switch (result.problem.type) {
        case types.unauthenticated: return 'Sesi perlu diperiksa. Pilih pemulihan sesi atau masuk kembali; permintaan sebelumnya tidak dikirim ulang otomatis.';
        case types.forbidden: return 'Permintaan tidak diizinkan. Periksa sesi sebelum mencoba secara eksplisit.';
        case types.invalid: return 'Bentuk permintaan tidak dapat diproses. Tidak ada pengiriman ulang otomatis.';
        case types.validation: return 'Periksa nama organisasi: 1–200 karakter, tanpa karakter kontrol.';
        case types.notObserved: return 'Hasil permintaan belum teramati. Ini bukan bukti pembatalan. Periksa organisasi Anda.';
        case types.expired: return 'Jendela pengiriman ulang berakhir. Ini bukan bukti permintaan sebelumnya dibatalkan.';
        case types.mismatch: return 'Identitas permintaan tidak cocok dengan nama semula. Jangan mengganti nama untuk permintaan ini.';
        case types.rate: return `Terlalu banyak permintaan. Tunggu ${result.wait ?? 60} detik sebelum mencoba secara eksplisit dengan identitas yang sama.`;
        case types.proof: return 'Verifikasi akun yang sama melalui Google diperlukan sebelum pratinjau baru.';
        case types.scope: return 'Daftar organisasi berubah. Tinjau seluruh cakupan lagi sebelum konfirmasi.';
        case types.preview: return 'Pratinjau tidak berlaku lagi. Minta pratinjau baru; belum ada konfirmasi penghapusan dari respons ini.';
        case types.manual: return 'Konteks ini memerlukan penilaian pengelola privasi. Jangan menutup atau memindahkan organisasi secara otomatis.';
        case types.pending: return 'Niat penghapusan diterima; pemrosesan belum terbukti selesai. Akses dapat ditahan. Pembatalan otomatis tidak tersedia.';
        case types.resource: return 'Organisasi tidak tersedia untuk akun ini. Ini bukan bukti pembatalan permintaan sebelumnya.';
        default: return `${context}: hasil belum dapat dipastikan. Periksa melalui langkah yang didukung; jangan mengirim permintaan baru otomatis.`;
    }
}
