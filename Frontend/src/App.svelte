<script>
    import { onMount } from 'svelte';
    import axios from 'axios';

    let medicines = [];
    let form = {
        name: '',
        ingredient: '',
        brand: '',
        expiry: '',
        quantity: 0
    };

    onMount(fetchMedicines);

    async function fetchMedicines() {
        try {
            const response = await axios.get('http://localhost:8080/medicines');
            medicines = response.data;
        } catch (error) {
            console.error("Error fetching medicines:", error);
        }
    }

    async function addMedicine() {
        try {
            await axios.post('http://localhost:8080/medicines', {
                nombre: form.name,
                ingrediente_activo: form.ingredient,
                marca: form.brand,
                fecha_vencimiento: form.expiry,
                cantidad: form.quantity
            });
            form = { name: '', ingredient: '', brand: '', expiry: '', quantity: 0 };
            fetchMedicines();
        } catch (error) {
            console.error("Error adding medicine:", error);
        }
    }

    async function deleteMedicine(id) {
        try {
            await axios.delete(`http://localhost:8080/medicines/${id}`);
            fetchMedicines();
        } catch (error) {
            console.error("Error deleting medicine:", error);
        }
    }
</script>

<style>
    body {
        font-family: Arial, sans-serif;
        background-color: #f4f4f4;
        margin: 0;
        padding: 0;
    }

    h1 {
        background-color: #333;
        color: #fff;
        padding: 1rem;
        text-align: center;
        margin: 0;
    }

    #medicine-form, #medicine-list {
        margin: 2rem;
        padding: 1rem;
        background-color: #fff;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    }

    label {
        display: block;
        margin: 0.5rem 0 0.2rem;
    }

    input, button {
        width: 100%;
        padding: 0.5rem;
        margin-bottom: 1rem;
    }

    table {
        width: 100%;
        border-collapse: collapse;
        margin-top: 1rem;
    }

    th, td {
        padding: 0.5rem;
        border: 1px solid #ccc;
        text-align: left;
    }

    button {
        background-color: #333;
        color: #fff;
        border: none;
        cursor: pointer;
    }

    button:hover {
        background-color: #555;
    }
</style>

<h1>Medicine Inventory</h1>
<div id="medicine-form">
    <h2>Add Medicine</h2>
    <form on:submit|preventDefault={addMedicine}>
        <label for="name">Name:</label>
        <input type="text" id="name" bind:value={form.name} required>
        <label for="ingredient">Ingredient:</label>
        <input type="text" id="ingredient" bind:value={form.ingredient} required>
        <label for="brand">Brand:</label>
        <input type="text" id="brand" bind:value={form.brand} required>
        <label for="expiry">Expiry Date:</label>
        <input type="date" id="expiry" bind:value={form.expiry} required>
        <label for="quantity">Quantity:</label>
        <input type="number" id="quantity" bind:value={form.quantity} required>
        <button type="submit">Add Medicine</button>
    </form>
</div>
<div id="medicine-list">
    <h2>Medicine List</h2>
    <table>
        <thead>
            <tr>
                <th>Name</th>
                <th>Ingredient</th>
                <th>Brand</th>
                <th>Expiry</th>
                <th>Quantity</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            {#each medicines as medicine}
                <tr>
                    <td>{medicine.nombre}</td>
                    <td>{medicine.ingrediente_activo}</td>
                    <td>{medicine.marca}</td>
                    <td>{medicine.fecha_vencimiento}</td>
                    <td>{medicine.cantidad}</td>
                    <td>
                        <button on:click={() => deleteMedicine(medicine.id)}>Delete</button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
