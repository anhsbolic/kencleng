export default function Home() {
  return (
    <>
      <header className="site-header" id="awal">
        <div className="page-width header-inner">
          <p className="site-name">Kencleng</p>
          <nav aria-label="Navigasi utama">
            <a className="header-link" href="#kejelasan">
              Tentang kejelasan
            </a>
          </nav>
        </div>
      </header>

      <main>
        <section className="hero page-width" aria-labelledby="hero-title">
          <div className="hero-main">
            <p className="eyebrow">Ruang untuk memberi dengan jernih</p>
            <h1 id="hero-title">
              Harapan tumbuh dari <em>hal yang jelas.</em>
            </h1>
            <p className="hero-intro">
              Kencleng adalah ruang penggalangan dana yang menempatkan cerita
              dan informasi berdampingan. Sebelum melangkah, pahami apa yang
              diketahui, apa yang belum, dan bagaimana perkembangan disampaikan.
            </p>
            <a className="primary-link" href="#kejelasan">
              Lihat cara kami menjelaskan
              <span aria-hidden="true">↗</span>
            </a>
          </div>
          <aside className="hero-note" aria-label="Catatan tentang pendekatan Kencleng">
            <span className="note-number">01 / Titik awal</span>
            <p>Cerita mengundang perhatian. Informasi memberi pegangan.</p>
          </aside>
        </section>

        <section className="clarity" id="kejelasan" aria-labelledby="clarity-title">
          <div className="page-width clarity-inner">
            <div className="clarity-heading">
              <p className="eyebrow">Kejelasan sebelum keputusan</p>
              <h2 id="clarity-title">Ruang untuk melihat lebih dekat.</h2>
              <p>
                Sebuah ajakan untuk memberi perlu disertai konteks yang bisa
                dipahami. Inilah cara kami memandang informasi dalam pengalaman
                penggalangan dana.
              </p>
            </div>
            <div className="clarity-points">
              <div className="clarity-point">
                <span aria-hidden="true">01</span>
                <div>
                  <h3>Tujuan yang dapat dipahami</h3>
                  <p>Mulai dari maksud penggalangan dan siapa yang menyampaikannya.</p>
                </div>
              </div>
              <div className="clarity-point">
                <span aria-hidden="true">02</span>
                <div>
                  <h3>Perkembangan dalam konteks</h3>
                  <p>Bedakan dana yang terkumpul dari pelaksanaan dan hasil yang dilaporkan.</p>
                </div>
              </div>
              <div className="clarity-point">
                <span aria-hidden="true">03</span>
                <div>
                  <h3>Yang belum diketahui tetap terlihat</h3>
                  <p>Informasi yang belum tersedia tidak perlu diisi dengan kepastian semu.</p>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section className="closing page-width" aria-labelledby="closing-title">
          <p className="eyebrow">Kencleng</p>
          <h2 id="closing-title">Mulai dari yang jelas.</h2>
          <a href="#awal">Kembali ke awal <span aria-hidden="true">↑</span></a>
        </section>
      </main>
    </>
  );
}
