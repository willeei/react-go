# CRUD de Bíblia com Go (SQLite) + React

Este projeto é um CRUD completo para cadastro de Bíblias, com:
- **Backend:** Go + SQLite
- **Frontend:** React + Axios

---

## ✨ Funcionalidades
- Cadastrar bíblia
- Listar todas as bíblias
- Editar uma bíblia
- Remover uma bíblia

---

## ☑️ Requisitos

### Go (Golang)
#### Windows:
1. Acesse: https://go.dev/dl/
2. Baixe e instale a versão estável

#### macOS:
```bash
brew install go
```

#### Linux (Ubuntu/Debian):
```bash
sudo apt update && sudo apt install -y golang
```

### Node.js + NPM com NVM
#### Windows:
1. Use o NVM para Windows: https://github.com/coreybutler/nvm-windows/releases
2. Instale Node.js:
```bash
nvm install 18
nvm use 18
```

#### macOS / Linux:
```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
source ~/.bashrc  # ou ~/.zshrc
nvm install 18
nvm use 18
```

---

## 📁 Estrutura do Projeto
```
biblia-crud/
├── backend/
│   ├── main.go
├── frontend/
│   └── (React App)
├── README.md
```

---

## 🚀 Backend com Go + SQLite

### Inicializar o projeto
```bash
cd backend
go mod init github.com/seunome/biblia-crud
```

### Instalar dependência SQLite
```bash
go get github.com/mattn/go-sqlite3
```

### main.go
```go
package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    _ "github.com/mattn/go-sqlite3"
    "github.com/gorilla/mux"
)

type Biblia struct {
    ID     int    `json:"id"`
    Nome   string `json:"nome"`
    Versao string `json:"versao"`
    Idioma string `json:"idioma"`
}

var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("sqlite3", "biblia.db")
    if err != nil {
        log.Fatal(err)
    }

    db.Exec("CREATE TABLE IF NOT EXISTS biblias (id INTEGER PRIMARY KEY AUTOINCREMENT, nome TEXT, versao TEXT, idioma TEXT)")

    r := mux.NewRouter()
    r.HandleFunc("/biblias", getBiblias).Methods("GET")
    r.HandleFunc("/biblias", createBiblia).Methods("POST")
    r.HandleFunc("/biblias/{id}", updateBiblia).Methods("PUT")
    r.HandleFunc("/biblias/{id}", deleteBiblia).Methods("DELETE")

    log.Println("Servidor iniciado em http://localhost:8080")
    http.ListenAndServe(":8080", r)
}

func getBiblias(w http.ResponseWriter, r *http.Request) {
    rows, _ := db.Query("SELECT * FROM biblias")
    var biblias []Biblia
    for rows.Next() {
        var b Biblia
        rows.Scan(&b.ID, &b.Nome, &b.Versao, &b.Idioma)
        biblias = append(biblias, b)
    }
    json.NewEncoder(w).Encode(biblias)
}

func createBiblia(w http.ResponseWriter, r *http.Request) {
    var b Biblia
    json.NewDecoder(r.Body).Decode(&b)
    stmt, _ := db.Prepare("INSERT INTO biblias(nome, versao, idioma) VALUES (?, ?, ?)")
    stmt.Exec(b.Nome, b.Versao, b.Idioma)
    w.WriteHeader(http.StatusCreated)
}

func updateBiblia(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var b Biblia
    json.NewDecoder(r.Body).Decode(&b)
    stmt, _ := db.Prepare("UPDATE biblias SET nome=?, versao=?, idioma=? WHERE id=?")
    stmt.Exec(b.Nome, b.Versao, b.Idioma, id)
}

func deleteBiblia(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    stmt, _ := db.Prepare("DELETE FROM biblias WHERE id=?")
    stmt.Exec(id)
}
```

---

## 💻 Frontend com React

### Criar o projeto React
```bash
npx create-react-app frontend
cd frontend
npm install axios
```

### src/App.jsx
```jsx
import React, { useEffect, useState } from 'react';
import axios from 'axios';

function App() {
  const [biblias, setBiblias] = useState([]);
  const [form, setForm] = useState({ nome: '', versao: '', idioma: '' });
  const [editId, setEditId] = useState(null);

  useEffect(() => {
    axios.get('http://localhost:8080/biblias').then(res => setBiblias(res.data));
  }, []);

  const submit = () => {
    if (editId) {
      axios.put(`http://localhost:8080/biblias/${editId}`, form).then(() => window.location.reload());
    } else {
      axios.post('http://localhost:8080/biblias', form).then(() => window.location.reload());
    }
  };

  const remove = (id) => {
    axios.delete(`http://localhost:8080/biblias/${id}`).then(() => window.location.reload());
  };

  return (
    <div className="App">
      <h1>Cadastro de Bíblia</h1>
      <input placeholder="Nome" onChange={e => setForm({ ...form, nome: e.target.value })} value={form.nome} />
      <input placeholder="Versão" onChange={e => setForm({ ...form, versao: e.target.value })} value={form.versao} />
      <input placeholder="Idioma" onChange={e => setForm({ ...form, idioma: e.target.value })} value={form.idioma} />
      <button onClick={submit}>{editId ? 'Atualizar' : 'Salvar'}</button>
      <ul>
        {biblias.map(b => (
          <li key={b.id}>
            {b.nome} ({b.versao}) - {b.idioma}
            <button onClick={() => { setForm(b); setEditId(b.id); }}>Editar</button>
            <button onClick={() => remove(b.id)}>Remover</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
```

---

## ▶️ Executar Projeto

### Backend
```bash
cd backend
go run main.go
```

### Frontend
```bash
cd frontend
npm start
```

Acesse:
- Frontend: http://localhost:3000
- API: http://localhost:8080/biblias

---

## 📄 Observações
- CORS pode precisar ser liberado no backend para uso real
- Este é um exemplo simples e didático
- Pode ser expandido com autenticação, validação, paginação, etc.

---

Feito com ❤️ por ChatGPT

