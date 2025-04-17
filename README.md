# 🧠 PsychoClient – Custom HTTP/2 Client with TLS Fingerprinting

**PsychoClient** is a custom-built HTTP client designed to work with **HTTP/2** and **custom TLS ClientHello specs** — originally created to help bypass advanced bot protection and emulate real browsers.

---

## ⚙️ What It Does

- 🧬 **Prebuilt TLS ClientHello Specs**  
  - Chrome, Firefox, iOS fingerprints (❗ these are outdated — you’ll need to update them in order to get it to work)
  
- 🌐 **HTTP Proxy Support**  
  - Full support for authenticated proxies (`http://ip:port:username:password`)
  
- 🍪 **Custom Cookie Jar**  
  - Handles request/response cookies manually
  
- 🔁 **Custom Sessions**  
  - Reuse connections, preserve cookies and headers
  
- 🧵 **Session Pooling**  
  - Reuse session objects via a pool (especially useful for multi-proxy environments like botting)  
  - The client will auto retry/reconnect if a session dies

---

## ❗ Disclaimer

- 🧪 **This is not plug-and-play. It’s outdated.**  
  You'll need to:
  - Update the TLS ClientHello specs
  - Some trouble shooting yourself

---

## 📜 License

This project is released under the **MIT License**.  
You can do **whatever the fuck you want** with it.

---

## 🧘 Sidenote

I still use this from time to time — though I don’t really mess with reverse engineering or botting these days. If you're into that space, this might be a good starting point.

---

### — Sicko
