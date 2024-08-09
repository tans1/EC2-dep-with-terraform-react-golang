import "./App.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Auth from "./pages/auth";
import Home from "./pages/home";
import { Navigate, Outlet } from "react-router-dom";
import { validateToken } from "./api/_auth";
import { useState, useEffect } from "react";

const ProtectedRoute = () => {
  const [authStatus, setAuthStatus] = useState<null | boolean>(null);
  console.log("inside the protected")
  useEffect(() => {
    const checkAuth = async () => {
      try {
        const result = await validateToken();
        setAuthStatus(result === 200);
      } catch (error) {
        setAuthStatus(false);
      }
    };

    checkAuth();
  }, []);

  if (authStatus === null) {
    return <div>Loading...</div>;
  }
  if (!authStatus) {
    return <Navigate to="/auth" replace />;
  }
  return <Outlet />;
};

const router = createBrowserRouter([
  {
    path: "/auth",
    element: <Auth />
  },
  {
    path: "/",
    element: <ProtectedRoute />,
    children: [
      {
        path: "/",
        element: <Home />,
        index: true
      }
    ]
  }
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
