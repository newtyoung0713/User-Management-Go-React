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

        const response = await axios.get("/users/profile", {
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

  const handleLogout = async () => {
    const token = localStorage.getItem("token");
    try {
      // Call logout endpoint
      await axios.post(
        "/users/logout",
        {},
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      // Clear token from localStorage
      localStorage.removeItem("token");

      // Redirect to login page
      navigate("/login");
    } catch (err) {
      console.error("Error during logout", err);
      // Even if the logout request fails, we still want to clear the token and redirect
      localStorage.removeItem("token");
      navigate("/login");
    }
  };

  return (
    <div className="flex flex-col items-center p-8">
      <h2 className="text-2xl font-semibold mb-4">User Profile</h2>
      {userData ? (
        <div className="w-full max-w-md p-4 rounded-lg shadow-lg">
          <h3 className="text-lg font-semibold">Welcome, {userData.username}</h3>
          <p>Email: {userData.email}</p>
          <p>Created At: {userData.created_at.toLocaleString()}</p>
          <p>Updated At: {userData.updated_at.toLocaleString()}</p>
          <button
            onClick={handleLogout}
            className="mt-4 px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600 transition-colors"
          >
            Logout
          </button>
        </div>
      ) : (
        <p>Loading user data...</p>
      )}
    </div>
  );
}
