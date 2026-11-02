import { cleanup, render, screen } from '@testing-library/react';
import { afterEach, it, expect } from 'vitest';
import Privacy from '../../app/privacy/page';
afterEach(cleanup);
it('presents full draft notice with semantic headings/tables and keeps writer metadata out of product flow', async () => { render(await Privacy()); expect(screen.getByRole('heading', { name: 'Pemberitahuan pemrosesan data' })).toBeVisible(); expect(screen.getByRole('heading', { name: 'Lama penyimpanan' })).toBeVisible(); expect(screen.getAllByRole('table')).toHaveLength(2); expect(screen.getByRole('heading', { name: 'Permintaan dan penghapusan akun' })).toBeVisible(); expect(screen.queryByText('Referensi untuk penulis dan reviewer')).toBeNull(); });
