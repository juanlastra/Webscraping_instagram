"use client";

import { useState } from "react";

interface ProfileData {
  usuario: string;
  posts: number;
  seguidores: number;
  seguidos: number;
}

export default function Home() {
  const [profiles, setProfiles] = useState<ProfileData[]>([]);
  const [inputLink, setInputLink] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!inputLink) return;

    try {
      const res = await fetch("http://localhost:8080/get-profile-data", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ link: inputLink }),
      });

      if (!res.ok) {
        throw new Error("Error en la solicitud");
      }

      const data: ProfileData = await res.json();

      // Añadir nuevo perfil a la tabla
      setProfiles((prev) => [...prev, data]);
      setInputLink(""); // Limpiar input
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="p-6 max-w-4xl mx-auto">
      <h1 className="text-2xl font-bold mb-6">Instagram Scraper</h1>

      <form onSubmit={handleSubmit} className="flex gap-4 mb-8">
        <input
          type="text"
          placeholder="Pega el link del perfil"
          value={inputLink}
          onChange={(e) => setInputLink(e.target.value)}
          className="border p-2 flex-1 rounded"
        />
        <button
          type="submit"
          className="bg-blue-600 text-white p-2 rounded hover:bg-blue-700"
        >
          Añadir
        </button>
      </form>

      <table className="w-full border-collapse">
        <thead>
          <tr className="bg-gray-200">
            <th className="border p-2 text-black">Usuario</th>
            <th className="border p-2 text-black">Posts</th>
            <th className="border p-2 text-black">Seguidores</th>
            <th className="border p-2 text-black">Seguidos</th>
          </tr>
        </thead>
        <tbody>
          {profiles.map((profile, index) => (
            <tr key={index} className="text-center">
              <td className="border p-2">{profile.usuario}</td>
              <td className="border p-2">{profile.posts}</td>
              <td className="border p-2">{profile.seguidores}</td>
              <td className="border p-2">{profile.seguidos}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
