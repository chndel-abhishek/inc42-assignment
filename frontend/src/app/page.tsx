import Head from 'next/head';
import axios from 'axios';
import { useState, useEffect } from 'react'; // Import useState and useEffect

interface User {
  id: number;
  name: string;
}

const HomePage = () => {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    axios.get('http://localhost:8085/api/users')
      .then(response => {
        setUsers(response.data);
      })
      .catch(error => {
        console.error(error);
      });
  }, []);

  return (
    <div>
      <Head>
        <title>My App</title>
      </Head>
      <h1>Users</h1>
      <ul>
        {users.map(user => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
    </div>
  );
};

export default HomePage;