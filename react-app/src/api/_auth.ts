import { apiInstance } from "./_index";

interface LoginData {
  username: string;
  password: string;
}
interface SignUpData {
  email: string;
  username: string;
  firstName: string;
  lastName: string;
  password: string;
  phone: string;
}

export function login({ username, password }: LoginData) {
  return apiInstance.post("/v1/auth/login", {
    username,
    password
  }).then((res) => {
    return res.data;
  })
  .catch((err) => {
    return err.response;
  })
}

export function register({
  email,
  username,
  firstName,
  lastName,
  password,
  phone
}: SignUpData) {
  return apiInstance.post("/v1/auth/register", {
    email,
    username,
    firstName,
    lastName,
    password,
    phone
  })
    .then((res) => {
      return res.data;
    })
    .catch((err) => {
      return err.response;
    })
  
  
  ;
}

export function logout() {
  return localStorage.removeItem("token");
}

export function getUser() {
  return apiInstance.get("/user");
}

export async function validateToken(): Promise<number> {
  return apiInstance
    .get("/v1/auth/validate")
    .then((res) => {
      return res.status;
    })
    .catch((err) => {
      return err.response.status;
    });
}
