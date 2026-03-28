import { useEffect, useState } from 'react';
import api from '../api';
import { LogOut, PackagePlus } from 'lucide-react';

export default function Dashboard() {
    const [products, setProducts] = useState([]);
    const [newProd, setNewProd] = useState({ name: '', price: 0 });
    const role = localStorage.getItem('role');

    const fetchProducts = async () => {
        const { data } = await api.get('/products/');
        setProducts(data);
    };

    const addProduct = async (e) => {
        e.preventDefault();
        try {
            await api.post('/products/', { ...newProd, price: parseFloat(newProd.price) });
            fetchProducts();
            alert("Product Added!");
        } catch (err) {
            alert("Unauthorized: Only admins can add products.");
        }
    };

    useEffect(() => { fetchProducts(); }, []);

    return (
        <div className="p-8">
            <div className="flex justify-between items-center mb-8">
                <h1 className="text-3xl font-bold">Product Catalog ({role})</h1>
                <button onClick={() => { localStorage.clear(); window.location.href = '/login'; }} className="flex items-center text-red-500"><LogOut size={18} className="mr-1" /> Logout</button>
            </div>

            {role === 'admin' && (
                <form onSubmit={addProduct} className="mb-8 p-4 border rounded-md bg-blue-50">
                    <h3 className="flex items-center font-bold mb-2"><PackagePlus size={18} className="mr-2" /> Add New Product (Admin Only)</h3>
                    <input className="border p-2 mr-2" placeholder="Product Name" onChange={e => setNewProd({ ...newProd, name: e.target.value })} />
                    <input type="number" className="border p-2 mr-2" placeholder="Price" onChange={e => setNewProd({ ...newProd, price: e.target.value })} />
                    <button className="bg-blue-600 text-white px-4 py-2 rounded">Save</button>
                </form>
            )}

            <div className="grid grid-cols-3 gap-4">
                {products?.map(p => (
                    <div key={p.id} className="border p-4 rounded shadow-sm">
                        <h4 className="font-bold">{p.name}</h4>
                        <p className="text-green-600">${p.price}</p>
                    </div>
                ))}
            </div>
        </div>
    );
}