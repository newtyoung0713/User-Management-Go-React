import axios from "axios";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

export function Login() {
  const [formData, setFormData] = useState({
    email: "",
    password: ""
  });
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prevData => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const { email, password } = formData;
    
    try {
      const response = await axios.post("http://localhost:1323/login", formData);
      localStorage.setItem("token", response.data.token);
      console.log("Login successfully:", formData);
      navigate("/profile");
    } catch (err) {
      console.error("Login failed: ", err);
      setFormData({
        email: "",
        password: "",
      });
      setError("Email or Password is wrong. Please confirm again");
    }

    if (!email || !password) {
      setError("All the fields are required.");
      return;
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h2 className="text-2xl font-semibold mb-4">Login</h2>
      <form className="w-80 space-y-4" onSubmit={handleSubmit}>
        <label className="flex justify-start">Email:</label>
        <input
          type="email"
          name="email"
          className="w-full border p-2"
          placeholder="example@me.com"
          value={formData.email}
          onChange={handleInputChange}
        />
        <label className="flex justify-start">Password:</label>
        <input
          type="password"
          name="password"
          className="w-full border p-2"
          placeholder="Password"
          value={formData.password}
          onChange={handleInputChange}
        />
        <button type="submit" className="w-full bg-blue-600 text-white p-2">
          Login
        </button>
        {error && <p className="text-red-500">{error}</p>}
      </form>
    </div>
  );
};
