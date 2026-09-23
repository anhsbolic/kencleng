# Patch Plan — CR-TOP-001

> Phase: Code Review / Patch plan  
> WORK_UNIT_ID: WU-S1-005  
> RUN_ID: CR-TOP-001  
> Author: Codex CLI agent (Reviewer)  
> Created: 2026-09-23  
> Model / Reasoning / Session: gpt-5.6-terra / medium / Fresh independent Code Review session  
> Target revision: 1b767c94732f95649299e33e8f4478c2c781f690  
> Workflow revision: not exposed by the invocation

## Objective

Memastikan `minio-init` tidak melaporkan sukses jika pembuatan bucket atau penyetelan salah satu anonymous policy gagal, sehingga R3 benar-benar dapat dikonvergensikan dan diverifikasi kemudian.

## Required patch

1. Ubah hanya `docker-compose.yml:30` dari `/bin/sh -c` menjadi `/bin/sh -ec` (alternatif setara: tambahkan `set -e;` sebagai perintah pertama skrip).
2. Jangan ubah urutan/rincian policy yang sudah benar: public tetap `download`; private tetap `none`; nama bucket, alias, credential, volume, ports, dan lifecycle service tidak berubah.
3. Jangan mengubah `until mc alias set ...; do sleep 1; done;`: dengan `-e`, kegagalan pada command kondisi `until` tetap digunakan untuk retry, sedangkan kegagalan setelah koneksi tersedia menghentikan initializer.

## Required targeted confirmation

- Periksa diff bahwa fail-fast berlaku pada seluruh perintah `mc mb` dan `mc anonymous set` di initializer.
- Jika runtime Podman Compose yang writable tersedia pada sesi patch, jalankan rendered config memakai runner repo (`podman-compose -f docker-compose.yml config` atau target Make ekuivalen) dan catat hasil. Jangan menyatakan policy persisted/runtime terverifikasi hanya dari rendered config.
- Testing tetap harus memeriksa policy persisted untuk bootstrap baru dan volume reuse serta evidence Caddy/media R1-R7 sesuai `TP-TOP-001` §12; patch ini tidak mengubah ownership tersebut.

## Non-goals

- Tidak menambahkan security headers/TLS Caddy.
- Tidak menambah Campaign endpoint, fixture, storage behavior, atau perubahan API/frontend/backend.
