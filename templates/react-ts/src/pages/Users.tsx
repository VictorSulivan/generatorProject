import React, { useState } from 'react';
import Button from '../components/common/Button';
import { useUsers, useCreateUser, useUpdateUser, useDeleteUser } from '../services/api';

const Users: React.FC = () => {
  const [newUser, setNewUser] = useState({ name: '', email: '' });
  const { data: users, isLoading } = useUsers();
  const createUser = useCreateUser();
  const updateUser = useUpdateUser();
  const deleteUser = useDeleteUser();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await createUser.mutateAsync(newUser);
    setNewUser({ name: '', email: '' });
  };

  if (isLoading) {
    return <div>Chargement...</div>;
  }

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-6">Gestion des Utilisateurs</h1>

      {/* Formulaire d'ajout */}
      <form onSubmit={handleSubmit} className="mb-8">
        <div className="flex gap-4">
          <input
            type="text"
            placeholder="Nom"
            value={newUser.name}
            onChange={(e) => setNewUser({ ...newUser, name: e.target.value })}
            className="border p-2 rounded"
          />
          <input
            type="email"
            placeholder="Email"
            value={newUser.email}
            onChange={(e) => setNewUser({ ...newUser, email: e.target.value })}
            className="border p-2 rounded"
          />
          <Button type="submit" isLoading={createUser.isPending}>
            Ajouter
          </Button>
        </div>
      </form>

      {/* Liste des utilisateurs */}
      <div className="grid gap-4">
        {users?.map((user) => (
          <div key={user.id} className="border p-4 rounded flex justify-between items-center">
            <div>
              <h2 className="font-semibold">{user.name}</h2>
              <p className="text-gray-600">{user.email}</p>
            </div>
            <div className="flex gap-2">
              <Button
                variant="outline"
                onClick={() => updateUser.mutate({ ...user, name: `${user.name} (modifié)` })}
              >
                Modifier
              </Button>
              <Button
                variant="destructive"
                onClick={() => deleteUser.mutate(user.id)}
              >
                Supprimer
              </Button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Users; 