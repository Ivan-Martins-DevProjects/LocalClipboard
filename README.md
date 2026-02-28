## 📋 Local Copy & Paste via WebSocket

Uma ferramenta leve escrita em **Go** que permite enviar textos para a área de transferência (clipboard) de outra máquina em tempo real usando **WebSocket**.

Ideal para quem quer copiar algo no celular ou em outra máquina e colar instantaneamente no computador.

---

## 🚀 Visão Geral

O projeto funciona como um **servidor local de clipboard**:

- Um **server WebSocket** roda na máquina destino.
- Qualquer cliente pode enviar texto para esse server.
- Ao receber a mensagem, o server:
  - processa o payload
  - copia automaticamente o conteúdo para o clipboard da máquina

Atualmente o envio pode ser feito via **Postman** ou qualquer cliente WebSocket.

---

## ✨ Funcionalidades

- ✅ Servidor WebSocket local  
- ✅ Cópia automática para o clipboard  
- ✅ Baixo consumo de recursos (Go)  
- ✅ Builds para:
  - Windows
  - Linux
- 🚧 Cliente mobile em desenvolvimento  
- 🚧 Interface gráfica (planejado)

---

## 🧠 Como Funciona
### Fluxo
1. O servidor é executado na porta 5858 na máquina que receberá o conteúdo.  
2. Um cliente WebSocket envia um texto.  
3. Ao enviar o payload, o texto:
   - é recebido pelo server  
   - é escrito diretamente no clipboard local

## Binários
### Linux
```bash
./build/bin/LocalClipboard
```
### Windows
```bash
./build/windowss
