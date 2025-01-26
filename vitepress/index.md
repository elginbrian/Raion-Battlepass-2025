---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  image:
    src: /sticker.png
    alt: sticker
  name: "RAION BATTLEPASS 2025"
  tagline: "App Programmer"
  actions:
    - theme: brand
      text: --> Dokumentasi API
      link: https://raion-battlepass.elginbrian.com/docs
---

<br>

# Sebelum Kita Mulai, Ngobrol Dulu Yuk!

Halo para calon developer! 👋 Selamat datang di Raion Community! Sebelum kalian jadi bagian dari keluarga besar Raion, ada challenge kecil yang harus diselesaikan, tapi tenang aja, challengenya ngga susah kok! 😊

Yuk, kita cari tahu bareng apa itu **Raion Battlepass 2025** dan gimana cara ikutnya!

## ⚔ KETENTUAN BATTLEPASS

Di battlepass ini, kalian diminta untuk bikin aplikasi sosial media **sederhana**, yang penting **bisa jalan** dan sesuai dengan spesifikasi yang sudah kita tentukan. Ini dia yang perlu kalian siapin:

1. **Aplikasi Harus Pakai API Raion Battlepass.**  
   API ini tuh kayak "otak" dari aplikasi kalian, yang membantu aplikasi buat ambil data, bikin post baru, ngedit, atau ngehapus post.

2. **Menggunakan Salah Satu Teknologi Berikut.**

   - Jetpack Compose
   - Flutter
   - React Native

   Pilih aja terserah, gak ngaruh ke penilaian kok.

<br>

3. **Fitur yang Harus Ada di dalam Aplikasi:**

Layaknya sosial media beneran, aplikasi yang kalian buat nanti harus punya fitur-fitur berikut. Kita udah nyiapin API dari backend yang kita bikin, jadi tugas kalian sekarang adalah mengintegrasikan API tersebut ke dalam aplikasi yang kalian kembangin.

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
- **Update User Information:** Pengguna bisa ganti nama profil mereka.  
  **Endpoint:** `PUT /api/users`
- **Get User Details by ID:** Pengguna bisa lihat detail pengguna lain.  
   **Endpoint:** `GET /api/users/{id}`

:::

<br>

4. **Desainnya Ga Usah Aneh-aneh.**  
   Gak perlu bagus-bagus desainnya, yang penting aplikasi kalian bisa jalan dan gampang dipakai. Tapi kalau UI/UX kalian keren, bakal ada **bonus poin!**

5. **Bonus Poin** juga untuk kalian yang bisa:
   - Menyimpan data dari API ke database lokal (contoh: pakai Room di Android atau SQLite di Flutter/React Native).
   - Menulis kode yang mudah dibaca dan dimengerti orang lain.
   - Membuat **unit testing** (minimal 1 aja).

## 📲 TENTANG API

API itu seperti jembatan yang bikin aplikasi kalian bisa berkomunikasi dengan server. Dokumentasi API yang lebih lengkap bisa kalian lihat di sini:

[Dokumentasi API Raion Battlepass](https://raion-battlepass.elginbrian.com/docs)

::: details Mengenal API lebih lanjut 📚

### Apa Itu API?

API (Application Programming Interface) adalah sebuah sistem yang memungkinkan dua aplikasi atau sistem untuk saling berkomunikasi. Coba bayangkan API seperti sebuah **jembatan** yang menghubungkan aplikasi yang kamu gunakan di handphone (frontend) dengan sistem yang ada di server (backend). Jadi, API adalah cara aplikasi untuk "ngobrol" dengan server untuk meminta atau mengirim data.

Misalnya, ketika kamu membuka aplikasi sosial media, aplikasi tersebut akan meminta data (seperti postingan atau informasi pengguna) ke server melalui API, dan server akan memberikan jawabannya.

### Bagaimana Cara Kerja API?

1. **Frontend Mengirimkan Permintaan (Request)**:

   - Aplikasi frontend (misalnya aplikasi di handphone) mengirimkan permintaan ke server melalui API. Permintaan ini bisa berisi data yang ingin kamu ambil (misalnya, daftar postingan) atau data yang ingin kamu kirimkan (misalnya, membuat postingan baru).
   - Permintaan ini dilakukan menggunakan protokol yang disebut HTTP. Biasanya, ada berbagai jenis permintaan seperti `GET` (untuk mengambil data), `POST` (untuk mengirim data baru), `PUT` (untuk memperbarui data), dan `DELETE` (untuk menghapus data).

2. **Backend Memproses Permintaan**:

   - Server backend akan menerima permintaan tersebut dan memprosesnya.
   - Jika kamu meminta data, server akan mengambilnya dari database dan mengirimkan jawabannya.
   - Jika permintaan berisi data baru (misalnya, kamu membuat postingan baru), server akan menyimpan data tersebut di database.

3. **Frontend Menerima Respons (Response)**:

   - Setelah permintaan diproses, server akan mengirimkan respons yang berisi data yang kamu minta atau konfirmasi bahwa tindakan kamu (misalnya, membuat postingan baru) berhasil.
   - Data yang dikirimkan biasanya dalam bentuk **JSON**, format yang mudah dibaca oleh aplikasi frontend.

### Jenis-jenis API

- **REST API**: API yang paling umum digunakan. API ini menggunakan URL untuk mengakses data dan biasanya menggunakan metode HTTP seperti `GET`, `POST`, `PUT`, dan `DELETE`.
- **GraphQL**: API yang memungkinkan aplikasi untuk meminta hanya data yang dibutuhkan, jadi tidak ada data yang terbuang (over-fetching) atau data yang kurang (under-fetching).
- **SOAP API**: API yang lebih kompleks dan biasanya digunakan dalam aplikasi yang lebih besar atau sistem lama.

### Manfaat API

- **Penghubung Antar Aplikasi**: API memungkinkan aplikasi yang berbeda, bahkan yang dibangun dengan teknologi yang berbeda, untuk saling berkomunikasi.
- **Akses ke Data dan Layanan**: API memungkinkan aplikasi untuk mengakses data yang ada di server atau layanan eksternal, seperti data cuaca atau layanan pembayaran.
- **Skalabilitas dan Modularitas**: Dengan API, aplikasi bisa dibangun dengan komponen yang terpisah dan saling berhubungan, sehingga lebih mudah untuk dikembangkan dan dikelola.

### API dalam Aplikasi Mobile

Dalam aplikasi mobile, API digunakan untuk menghubungkan aplikasi di handphone dengan server. Beberapa contoh penggunaan API dalam aplikasi mobile:

- **Mengambil data**: Misalnya, aplikasi mengambil daftar postingan atau informasi pengguna.
- **Mengirim data**: Misalnya, aplikasi mengirimkan data login atau membuat postingan baru.
- **Autentikasi dan Otorisasi**: API juga digunakan untuk memverifikasi identitas pengguna, seperti saat login ke aplikasi.

Jadi, meskipun aplikasi frontend (seperti di handphone) dan backend (seperti server) dibangun dengan teknologi yang berbeda, API membantu keduanya untuk berkomunikasi dengan cara yang seragam.

:::

::: details Contoh kode yang sudah kita disiapkan di backend ⚙

Sebagai contoh, di backend kita udah nyiapin function buat menangani request untuk mengambil semua postingan (GetAllPosts). Function ini akan mengambil semua postingan dari database dan mengembalikannya dalam bentuk JSON.

```Go
func (h *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := h.postService.FetchAllPosts()
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	var postResponse []domain.PostResponse
	for _, post := range posts {
		postResponse = append(postResponse, domain.PostResponse{
			ID:        post.ID,
			UserID:    post.UserID,
			Caption:   post.Caption,
			ImageURL:  post.ImageURL,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response.Success(c, postResponse, fiber.StatusOK)
}
```

Function tersebut nantinya akan menampilkan data dari endpoint API https://raion-battlepass.elginbrian.com/api/posts. Contoh data yang ditampilkan bisa di lihat di bawah ini dimana data tersebut merupakan respons JSON yang berisi informasi mengenai semua postingan yang telah diunggah ke server.

<iframe 
    src="https://raion-battlepass.elginbrian.com/api/posts" 
    width="100%" 
    height="240px" 
    frameborder="0" 
    style="border: none; margin-top: 20px; border-radius: 8px; box-shadow: 0 4px 8px rgba(0,0,0,0.1);">
</iframe>

:::

::: details Contoh kode yang harus kalian buat di sisi apps-nya 📱

Terus di sisi apps nya kalian perlu mengirimkan request ke server untuk mengambil data tersebut. Kalau di kode yang kita contohin di bawah ini, kita menggunakan Retrofit (sebuah library di Android untuk membuat request HTTP) untuk request GET untuk mengambil daftar postingan dari server.

```kotlin
data class PostResponse(
    val id: String,
    val user_id: String,
    val caption: String,
    val image_url: String,
    val created_at: String,
    val updated_at: String
)

data class ApiResponse(
    val status: String,
    val data: List<PostResponse>
)

interface ApiService {
    @GET("api/posts")
    fun getAllPosts(): Call<ApiResponse>
}

private val retrofit = Retrofit.Builder()
    .baseUrl("https://raion-battlepass.elginbrian.com/")
    .addConverterFactory(GsonConverterFactory.create())
    .build()

private val apiService = retrofit.create(ApiService::class.java)

```

Tapi, kalian dibebasin kok mau pakai library apa untuk bikin requestnya, ngga harus pakai Retrofit. Ga akan ngaruh ke penilaian juga.

:::

## 📩 PENGUMPULAN KARYA

Kalau udah selesai, jangan lupa kirim karya terbaik kalian lewat form di bawah ini yaa! Good Luck and see you at Raion Community! 💪✨

[Form Pengumpulan Karya](https://forms.gle/2EuGEa7vDSDpJBQt9)

::: info Bentuk Pengumpulan

Link folder Google Drive yang berisi:

- File **.apk** aplikasi kalian 📱
- Video **demo aplikasi** (maks. 2 menit) 🎥
- Video **penjelasan kode program** (maks. 5 menit) 💻
- Link **Repository GitHub** (bisa dimasukkan ke file .txt / .pdf) 📄

  :::

::: warning Note
Jangan nge-update folder Google Drive atau repository GitHub kalian setelah melebihi deadline ya guys, nanti bisa-bisa diskualifikasi, hihihi... 😜😁

:::

## 📞 CONTACT PERSON!

Kalau ada pertanyaan atau butuh bantuan, jangan takut buat tanya ya!

- **Elgin**: [0857-4980-6571](https://wa.me/6285749806571)
- **Ovan**: [0812-1675-2495](https://wa.me/6281216752495)
