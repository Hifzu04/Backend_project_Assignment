import { useState } from 'react';
import api from '../api';
import { useNavigate } from 'react-router-dom';

export default function Auth({ mode = 'login' }) {
    const [formData, setFormData] = useState({ email: '', password: '', name: '', role: 'user' });
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const endpoint = mode === 'login' ? '/auth/login' : '/auth/register';
            const { data } = await api.post(endpoint, formData);

            if (mode === 'login') {
                localStorage.setItem('token', data.token);
                localStorage.setItem('role', data.role);
                navigate('/dashboard');
            } else {
                alert("Registered! Please login.");
                navigate('/login');
            }
        } catch (err) {
            setError(err.response?.data?.error || "Something went wrong");
        }
    };

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
            <form onSubmit={handleSubmit} className="p-8 bg-white shadow-md rounded-lg w-96">
                <h2 className="text-2xl font-bold mb-4 capitalize">{mode}</h2>
                {error && <p className="text-red-500 mb-2">{error}</p>}
                {mode === 'register' && (
                    <>
                        <input className="border p-2 w-full mb-2" placeholder="Name" onChange={e => setFormData({ ...formData, name: e.target.value })} />
                        <select className="border p-2 w-full mb-2" onChange={e => setFormData({ ...formData, role: e.target.value })}>
                            <option value="user">User</option>
                            <option value="admin">Admin</option>
                        </select>
                    </>
                )}
                <input className="border p-2 w-full mb-2" placeholder="Email" onChange={e => setFormData({ ...formData, email: e.target.value })} />
                <input type="password" className="border p-2 w-full mb-4" placeholder="Password" onChange={e => setFormData({ ...formData, password: e.target.value })} />
                <button className="bg-blue-600 text-white w-full py-2 rounded">{mode}</button>
            </form>
        </div>
    );
}