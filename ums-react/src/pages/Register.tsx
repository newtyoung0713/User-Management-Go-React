import axios from "axios";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

export function Register() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  });
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prevData => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await axios.post("http://localhost:1323/register", formData);
      localStorage.setItem("token", response.data.token);
      navigate("/login");
      console.log(response.data.message);
      console.log("Register successfully:", formData);
    } catch (err) {
      setError("Failure for register: " + err + " Please contact to our team.");
    }
    const { username, email, password, confirmPassword } = formData;
    
    if (!username || !email || !password || !confirmPassword) {
      setError("All the fields are required.");
      return;
    }

    if (password !== confirmPassword) {
      setError("Passwords do not match.");
      return;
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h2 className="text-2xl font-semibold mb-4">Register</h2>
      <form className="w-80 space-y-4" onSubmit={handleSubmit}>
        <label className="flex justify-start">Username:</label>
        <input
          type="text"
          name="username"
          className="w-full border p-2"
          placeholder="Username"
          value={formData.username}
          onChange={handleInputChange}
          required
        />
        <label className="flex justify-start">Email:</label>
        <input
          type="email"
          name="email"
          className="w-full border p-2"
          placeholder="Email"
          value={formData.email}
          onChange={handleInputChange}
          required
        />
        <label className="flex justify-start">Password:</label>
        <input
          type="password"
          name="password"
          className="w-full border p-2"
          placeholder="Password"
          value={formData.password}
          onChange={handleInputChange}
          required
        />
        <label className="flex justify-start">Confirm Password:</label>
        <input
          type="password"
          name="confirmPassword"
          className="w-full border p-2"
          placeholder="Confirm Password"
          value={formData.confirmPassword}
          onChange={handleInputChange}
          required
        />
        <button type="submit" className="w-full bg-blue-600 text-white p-2">
          Register
        </button>
        {error && <p className="text-red-500">{error}</p>}
      </form>
    </div>
  );
};
