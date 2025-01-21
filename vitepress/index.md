---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "RAION BATTLEPASS 2025 🦁"
  tagline: "Mobile Developer Division"
  actions:
    - theme: brand
      text: Dokumentasi API
      link: https://raion-battlepass.elginbrian.com/docs
    - theme: alt
      text: Contact Person
      link: https://wa.me/6285749806571
---

# Sebelum Mulai, Ngobrol Dulu Yok!

Halo teman-teman developer muda! 👋 Selamat datang di Raion Community! Sebelum kalian jadi bagian dari keluarga besar Raion, ada tantangan kecil yang harus diselesaikan, tapi tenang aja, kita bakal bantu kalian dengan cara yang gampang kok! 😊

Yuk, kita cari tahu bareng apa itu **Raion Battlepass 2025** dan gimana cara ikutnya!

## ⚖ KETENTUAN BATTLEPASS

Di battlepass ini, kalian diminta untuk bikin aplikasi **sederhana**, yang penting **bisa jalan** dan sesuai dengan spesifikasi yang kita tentukan. Ini dia yang perlu kalian siapin:

1. **Aplikasi Harus Pakai API Raion Battlepass.**  
   API ini tuh kayak "otak" dari aplikasi kalian, yang membantu aplikasi buat ambil data, bikin post baru, ngedit, atau ngehapus post.

2. **Menggunakan Salah Satu Teknologi Berikut.**

   - Jetpack Compose
   - Flutter
   - React Native

   Pilih aja terserah, gak ngaruh ke penilaian kok.

<br>

3. **Fitur yang Harus Ada di dalam Aplikasi:**

::: details Fitur Autentikasi 🔐

- **Register:** Pengguna bisa bikin akun baru.  
  **Endpoint:** `POST /api/auth/register`
- **Login:** Pengguna bisa login ke aplikasi.  
  **Endpoint:** `POST /api/auth/login`
- **Change Password:** Pengguna bisa ganti password mereka.  
  **Endpoint:** `PUT /api/auth/change-password`
- **Get Current User:** Menampilkan informasi pengguna yang sedang login.  
   **Endpoint:** `GET /api/auth/current-user`

:::

::: details Fitur Postingan 📷

- **Create Post:** Pengguna bisa buat post baru.  
  **Endpoint:** `POST /api/posts`
- **Read Posts:** Pengguna bisa lihat daftar semua post.  
  **Endpoint:** `GET /api/posts`
- **Read Post by ID:** Pengguna bisa lihat post tertentu berdasarkan ID.  
  **Endpoint:** `GET /api/posts/{id}`
- **Update Post:** Pengguna bisa edit post yang sudah dibuat.  
  **Endpoint:** `PUT /api/posts/{id}`
- **Delete Post:** Pengguna bisa hapus post tertentu.  
  **Endpoint:** `DELETE /api/posts/{id}`
- **Get Posts by User ID:** Pengguna bisa lihat post milik user tertentu.  
   **Endpoint:** `GET /api/posts/user/{user_id}`

:::

::: details Fitur Searching 🔍

- **Search Posts:** Pengguna bisa cari post pakai kata kunci tertentu.  
  **Endpoint:** `GET /api/search/posts`
- **Search Users:** Pengguna bisa cari teman lain berdasarkan nama atau username.  
  **Endpoint:** `GET /api/search/users`

:::

::: details Fitur Manajemen User 👩‍🦰

- **Get All Users:** Pengguna bisa lihat daftar semua orang di aplikasi.  
  **Endpoint:** `GET /api/users`
- **Update User Information:** Pengguna bisa ganti nama, email, atau foto profil mereka.  
  **Endpoint:** `PUT /api/users`
- **Get User Details by ID:** Pengguna bisa lihat detail pengguna lain.  
   **Endpoint:** `GET /api/users/{id}`

:::

<br>

4. **Desainnya Ga Usah Aneh-aneh.**  
   Gak perlu bagus-bagus desainnya, yang penting aplikasi kalian bisa jalan dan gampang dipakai. Tapi kalau UI/UX kalian keren, ada **bonus poin! 🌟**

5. **Bonus Poin** untuk kalian yang bisa:
   - Menyimpan data dari API ke database lokal (contoh: pakai Room di Android atau SQLite di Flutter/React Native).
   - Menulis kode yang mudah dibaca dan dimengerti orang lain.
   - Membuat **unit testing** (minimal 1 aja).

## 📲 TENTANG API

API itu seperti jembatan yang bikin aplikasi kalian bisa berkomunikasi dengan server.  
Dokumentasi API yang lebih lengkap bisa kalian lihat di sini:  
[Dokumentasi API Raion Battlepass](https://raion-battlepass.elginbrian.com/docs)

## 📩 CARA KIRIM HASIL KARYA KAMU

> **Tenang aja**, gak ada aturan yang ribet kok, yang penting hasil akhirnya adalah **aplikasi mobile**.  
> Kalau sudah selesai, kirim karya terbaik kalian lewat form ini:

[Form Pengumpulan Karya](#) _(link form menyusul ya!)_

## 📞 CONTACT PERSON!

Kalau ada pertanyaan atau butuh bantuan, jangan takut buat tanya ya!

- **Elgin**: [0857-4980-6571](https://wa.me/6285749806571)
- **Ovan**: (Kontak menyusul)

## 🔗 LINK PENTING

- [Github Battlepass Repository](https://github.com/elginbrian/Raion-Battlepass-2025)
