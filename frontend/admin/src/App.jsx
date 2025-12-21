import { useEffect, useState } from "react";

const API = "http://localhost:8080/api/v1";

export default function AdminApp() {
  const [tab, setTab] = useState("users");

  return (
    <div className="p-8 max-w-5xl mx-auto">
      <h1 className="text-2xl font-bold mb-6">Admin Panel</h1>

      <div className="flex gap-4 mb-6">
        <button className="btn" onClick={() => setTab("users")}>Users</button>
        <button className="btn" onClick={() => setTab("posts")}>Posts</button>
      </div>

      {tab === "users" && <Users />}
      {tab === "posts" && <Posts />}
    </div>
  );
}

function Users() {
  const [users, setUsers] = useState([]);
  const [form, setForm] = useState({ username: "", email: "", password: "" });
  const [editing, setEditing] = useState(null);

  const load = async () => {
    const res = await fetch(`${API}/users`);
    setUsers(await res.json());
  };

  useEffect(() => { load(); }, []);

  const create = async () => {
    await fetch(`${API}/users`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(form),
    });
    setForm({ username: "", email: "", password: "" });
    load();
  };

  const startEdit = (u) => {
    setEditing(u.id);
    setForm({ username: u.username, email: u.email, password: "" });
  };

  const update = async () => {
    await fetch(`${API}/users/${editing}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username: form.username,
        email: form.email,
        ...(form.password && { password: form.password }),
      }),
    });
    setEditing(null);
    setForm({ username: "", email: "", password: "" });
    load();
  };

  return (
    <div>
      <h2 className="text-xl font-semibold">Users</h2>

      <div className="flex gap-2 my-4">
        <input placeholder="username" value={form.username} onChange={e => setForm({ ...form, username: e.target.value })} />
        <input placeholder="email" value={form.email} onChange={e => setForm({ ...form, email: e.target.value })} />
        <input
          type="password"
          placeholder={editing ? "new password (optional)" : "password"}
          value={form.password}
          onChange={e => setForm({ ...form, password: e.target.value })}
        />

        {editing ? (
          <>
            <button className="btn" onClick={update}>Update</button>
            <button className="btn" onClick={() => setEditing(null)}>Cancel</button>
          </>
        ) : (
          <button className="btn" onClick={create}>Create</button>
        )}
      </div>

      <table className="w-full border">
        <thead>
          <tr><th>ID</th><th>Username</th><th>Email</th><th></th></tr>
        </thead>
        <tbody>
          {users.map(u => (
            <tr key={u.id}>
              <td>{u.id}</td>
              <td>{u.username}</td>
              <td>{u.email}</td>
              <td><button className="btn" onClick={() => startEdit(u)}>Edit</button></td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

function Posts() {
  const [posts, setPosts] = useState([]);
  const [form, setForm] = useState({ title: "", content: "" });

  const load = async () => {
    const res = await fetch(`${API}/posts`);
    setPosts(await res.json());
  };

  useEffect(() => { load(); }, []);

  const create = async () => {
    await fetch(`${API}/posts`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(form),
    });
    setForm({ title: "", content: "" });
    load();
  };

  const update = async (id) => {
    const title = prompt("title");
    const content = prompt("content");
    await fetch(`${API}/posts/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title, content }),
    });
    load();
  };

  return (
    <div>
      <h2 className="text-xl font-semibold">Posts</h2>

      <div className="flex gap-2 my-4">
        <input placeholder="title" value={form.title} onChange={e => setForm({ ...form, title: e.target.value })} />
        <input placeholder="content" value={form.content} onChange={e => setForm({ ...form, content: e.target.value })} />
        <button className="btn" onClick={create}>Create</button>
      </div>

      <table className="w-full border">
        <thead>
          <tr><th>ID</th><th>Title</th><th></th></tr>
        </thead>
        <tbody>
          {posts.map(p => (
            <tr key={p.id}>
              <td>{p.id}</td>
              <td>{p.title}</td>
              <td><button className="btn" onClick={() => update(p.id)}>Edit</button></td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
