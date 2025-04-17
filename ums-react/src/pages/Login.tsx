import axios from "axios";
import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";

interface LoginProps {
  setIsAuthenticated: (value: boolean) => void;
}

export function Login({ setIsAuthenticated }: LoginProps) {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
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

    if (!email || !password) {
      setError("All the fields are required.");
      return;
    }

    try {
      const response = await axios.post("/users/login", formData);
      const { token } = response.data;

      // Save the token to local storage
      localStorage.setItem("token", token);

      // Set axios default headers
      axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;

      // Update authentication status
      setIsAuthenticated(true);

      console.log("Login successful");
      navigate("/profile");
    } catch (err) {
      console.error("Login failed: ", err);
      setFormData({
        email: "",
        password: "",
      });
      setError("Email or Password is incorrect. Please try again.");
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
        <p className="text-center text-sm">
          Don't have an account?{" "}
          <Link to="/register" className="text-blue-600 hover:text-blue-800">
            Register here
          </Link>
        </p>
        <button type="submit" className="w-full bg-blue-600 text-white p-2">
          Login
        </button>
        {error && <p className="text-red-500">{error}</p>}
      </form>
    </div>
  );
}
