import { useEffect, useState } from "react"
import "./App.css"
import axios from "axios";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Register } from "./pages/Register";
import { Login } from "./pages/Login";
import { UserProfile } from "./pages/UserProfile";

function App() {
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>("");

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get("http://localhost:1323/");
        console.log("Response data:", response.data);
      } catch (err) {
        console.error("Error fetching data:", err);
        setError("Error fetching data: ");
      } finally {
        setLoading(false);
      }
    };
  
    fetchData();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;
  
  return (
    <Router>
      <div>
        <h1 className="text-4xl font-bold text-blue-500">
          User Management System
        </h1>
        <Routes>
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/profile" element={<UserProfile />} />
        </Routes>
      </div>
    </Router>
  )
}

export default App
