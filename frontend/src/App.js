import React, { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [biblias, setBiblias] = useState([]); // nunca null
  const [form, setForm] = useState({ nome: "", versao: "", idioma: "" });
  const [editId, setEditId] = useState(null);

  useEffect(() => {
    axios
      .get("http://localhost:8080/biblias")
      .then((res) => {
        if (Array.isArray(res.data)) {
          setBiblias(res.data);
        } else {
          console.error("Resposta inesperada:", res.data);
          setBiblias([]);
        }
      })
      .catch((err) => {
        console.error("Erro ao buscar bíblias:", err.message);
        setBiblias([]); // garante que nunca seja null
      });
  }, []);

  const submit = () => {
    const action = editId
      ? axios.put(`http://localhost:8080/biblias/${editId}`, form)
      : axios.post("http://localhost:8080/biblias", form);

    action.then(() => {
      setForm({ nome: "", versao: "", idioma: "" });
      setEditId(null);
      window.location.reload();
    });
  };

  const remove = (id) => {
    axios.delete(`http://localhost:8080/biblias/${id}`).then(() => {
      window.location.reload();
    });
  };

  return (
    <div className="App">
      <h1>Cadastro de Bíblia</h1>
      <input
        placeholder="Nome"
        onChange={(e) => setForm({ ...form, nome: e.target.value })}
        value={form.nome}
      />
      <input
        placeholder="Versão"
        onChange={(e) => setForm({ ...form, versao: e.target.value })}
        value={form.versao}
      />
      <input
        placeholder="Idioma"
        onChange={(e) => setForm({ ...form, idioma: e.target.value })}
        value={form.idioma}
      />
      <button onClick={submit}>{editId ? "Atualizar" : "Salvar"}</button>

      <ul>
        {Array.isArray(biblias) &&
          biblias.map((b) => (
            <li key={b.id}>
              {b.nome} ({b.versao}) - {b.idioma}
              <button
                onClick={() => {
                  setForm({ nome: b.nome, versao: b.versao, idioma: b.idioma });
                  setEditId(b.id);
                }}
              >
                Editar
              </button>
              <button onClick={() => remove(b.id)}>Remover</button>
            </li>
          ))}
      </ul>
    </div>
  );
}

export default App;
