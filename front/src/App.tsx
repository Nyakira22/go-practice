import React from 'react';
import { useState,useEffect } from 'react';
import axios from "axios";
import logo from './logo.svg';
import './App.css';

type Fruit = {
	id: number;
	name: string;
	icon: string;
}

function App() {
	const [fruits, setFruits] = useState<Fruit[]>([{ id: 0, name: "", icon: "" }])
	useEffect(() => {
		(
			async () => {
				const data = await axios.get("http://localhost:8080")
				setFruits(data.data)
			}
		)()
	}, [])
	return (
	<div>
	{fruits.map(fruit => (
		<p key={fruit.id}>
			<span>{fruit.name}</span>
			<span>{fruit.icon}</span>
		</p>
	))}
    </div>

	);
}

export default App;
