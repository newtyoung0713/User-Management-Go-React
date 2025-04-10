import { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

interface User {
  id: number;
  username: string;
  email: string;
  password: string;
  created_at: Date;
  updated_at: Date;
  is_deleted: boolean;
}

export function UserProfile() {
  const [userData, setUserData] = useState<User | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchUserData = async () => {
      const token = localStorage.getItem("token");
      try {
        if (!token) {
          navigate("/login"); // Redirect to login if no token found
          return;
        }
        console.log("Token", token);
        const response = await axios.get("http://localhost:1323/profile", {
          headers: { Authorization: `Bearer ${token}` },
        });

        // Ensure 'created_at' and 'updated_at' are parsed as Date objects
        const data = response.data;
        data.created_at = new Date(data.created_at);
        data.updated_at = new Date(data.updated_at);
        
        setUserData(response.data);
      } catch (err) {
        console.error("Error fetching user data", err);
        navigate("/login"); // Redirect to login if error
      }
    };
    fetchUserData();
  }, [navigate]);

  return (
    <div className="flex flex-col items-center p-8 bg-gray-100">
      <h2 className="text-2xl font-semibold mb-4">User Profile</h2>
      {userData ? (
        <div className="w-full max-w-md p-4 bg-white rounded-lg shadow-lg">
          <h3 className="text-lg font-semibold">Welcome, {userData.username}</h3>
          <p>Email: {userData.email}</p>
          <p>Created At: {userData.created_at.toLocaleString()}</p>
          <p>Updated At: {userData.updated_at.toLocaleString()}</p>
        </div>
      ) : (
        <p>Loading user data...</p>
      )}
    </div>
  );
}
