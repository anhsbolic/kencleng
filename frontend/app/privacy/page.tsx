import Link from 'next/link';
import { readFile } from 'node:fs/promises';
import path from 'node:path';
function inline(text: string) { return text.split(/(\*\*[^*]+\*\*)/g).map((part, i) => part.startsWith('**') ? <strong key={i}>{part.slice(2, -2)}</strong> : part); }
function Notice({ text }: {
    text: string;
}) {
    const blocks = text.split('\n\n').filter(Boolean);
    return blocks.map((block, i) => {
        if (block.startsWith('### '))
            return <h2 key={i}>{block.slice(4)}</h2>;
        if (block.startsWith('|')) {
            const rows = block.split('\n').map(row => row.split('|').slice(1, -1).map(cell => cell.trim()));
            return <div className="table-scroll" key={i}><table><thead><tr>{rows[0].map((cell, j) => <th scope="col" key={j}>{inline(cell)}</th>)}</tr></thead><tbody>{rows.slice(2).map((row, j) => <tr key={j}>{row.map((cell, k) => <td key={k}>{inline(cell)}</td>)}</tr>)}</tbody></table></div>;
        }
        return <p key={i}>{inline(block)}</p>;
    });
}
export default async function Privacy() {
    const source = await readFile(path.resolve(process.cwd(), '../docs/product/outcomes/org-01/privacy-notice-draft.md'), 'utf8');
    const draft = source.split('## Draf pemberitahuan lengkap\n\n')[1].split('## Referensi untuk penulis dan reviewer')[0];
    return <><header className="site-header"><Link href="/" className="brand">Kencleng</Link><Link href="/organizations">Kembali ke organisasi</Link></header><main className="narrow"><div className="eyebrow">DRAFT · belum untuk production</div><h1>Pemberitahuan pemrosesan data</h1><p className="lead">Yayasan Kencleng Indonesia · care@kencleng.dev</p><p>Dokumen lengkap berikut adalah draf untuk pengembangan, bukan pemberitahuan final yang telah dipublikasikan. Penyedia, lokasi, dasar hukum akhir, operasional kanal, dan bukti penghapusan tetap menjadi gate sebelum data nyata atau production.</p><article className="notice-draft" aria-label="Draf lengkap pemberitahuan privasi"><Notice text={draft}/></article></main></>;
}
