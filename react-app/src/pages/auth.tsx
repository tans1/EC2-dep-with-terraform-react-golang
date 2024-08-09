import { useState } from "react";
import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
  CardFooter
} from "../components/ui/card";
import { Label } from "@radix-ui/react-label";
import { Input } from "../components/ui/input";
import { Button } from "../components/ui/button";
import { login, register } from "../api/_auth";
import { useNavigate } from "react-router-dom";

export default function Auth() {
  const [signUp, setSignUp] = useState(false);
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    email: "",
    username: "",
    firstName: "",
    lastName: "",
    password: "",
    confirmPassword: "",
    phone: ""
  });

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const handleChange = (e: { target: { id: any; value: any } }) => {
    const { id, value } = e.target;
    setFormData({ ...formData, [id]: value });
  };
  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    if (signUp) {
      register(formData)
        .then(() => {
          setSignUp(false);
          setFormData({
            email: "",
            username: "",
            firstName: "",
            lastName: "",
            password: "",
            confirmPassword: "",
            phone: ""
          });
        })
        .catch((err) => {
          console.log(err);
        });
    } else {
      login(formData)
        .then((res) => {
          localStorage.setItem("token", res.data.token);
          navigate("/", { replace: true });
        })
        .catch((err) => {
          console.log(err);
        });
    }
  };
  return (
    <div className="w-full h-screen flex justify-center items-center">
      {signUp ? (
        <Card className="w-[550px]">
          <form onSubmit={handleSubmit}>
            <CardHeader>
              <CardTitle>Sign Up</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid w-full items-center gap-4">
                <div className="flex flex-col items-start space-y-1.5">
                  <Label htmlFor="email">Email</Label>
                  <Input
                    id="email"
                    placeholder="Email"
                    value={formData.email}
                    onChange={handleChange}
                  />
                </div>
                <div className="flex flex-col items-start space-y-1.5">
                  <Label htmlFor="username">Username</Label>
                  <Input
                    id="username"
                    placeholder="Username"
                    value={formData.username}
                    onChange={handleChange}
                  />
                </div>
                <div className="flex flex-col items-start space-y-1.5">
                  <Label htmlFor="firstName">First Name</Label>
                  <Input
                    id="firstName"
                    placeholder="First Name"
                    value={formData.firstName}
                    onChange={handleChange}
                  />
                </div>
                <div className="flex flex-col items-start space-y-1.5">
                  <Label htmlFor="lastName">Last Name</Label>
                  <Input
                    id="lastName"
                    placeholder="Last Name"
                    value={formData.lastName}
                    onChange={handleChange}
                  />
                </div>
                <div className="flex flex-col items-start space-y-1.5">
                  <Label htmlFor="password">Password</Label>
                  <Input
                    id="password"
                    type="password"
                    placeholder="Password"
                    value={formData.password}
                    onChange={handleChange}
                  />
                </div>
                <div className="flex flex-col items-start space-y-1.5">
                  <Label htmlFor="confirmPassword">Confirm Password</Label>
                  <Input
                    id="confirmPassword"
                    type="password"
                    placeholder="Confirm Password"
                    value={formData.confirmPassword}
                    onChange={handleChange}
                  />
                </div>
                <div className="flex flex-col items-start space-y-1.5">
                  <Label htmlFor="phone">Phone</Label>
                  <Input
                    id="phone"
                    placeholder="Phone"
                    value={formData.phone}
                    onChange={handleChange}
                  />
                </div>
              </div>
            </CardContent>
            <CardFooter className="grid mt-5">
              <Button className="col-12" type="submit">
                Submit
              </Button>
              <p className="mt-2">
                Do you have an account?{" "}
                <Button variant="link" onClick={() => setSignUp(false)}>
                  Sign In
                </Button>
              </p>
            </CardFooter>
          </form>
        </Card>
      ) : (
        <Card className="w-[350px]">
          <form onSubmit={handleSubmit}>
            <CardHeader>
              <CardTitle>Sign In</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid w-full items-center gap-4">
                <div className="flex flex-col items-start space-y-1.5">
                  <Label htmlFor="username">Username</Label>
                  <Input
                    id="username"
                    placeholder="username"
                    value={formData.username}
                    onChange={handleChange}
                  />
                </div>
                <div className="flex flex-col items-start space-y-1.5">
                  <Label htmlFor="password">Password</Label>
                  <Input
                    id="password"
                    placeholder="Password"
                    value={formData.password}
                    onChange={handleChange}
                  />
                </div>
              </div>
            </CardContent>
            <CardFooter className="grid mt-5">
              <Button className="col-12" type="submit">
                Submit
              </Button>
              <p className="mt-2">
                Do you not have an account?{" "}
                <Button variant="link" onClick={() => setSignUp(true)}>
                  Sign Up
                </Button>
              </p>
            </CardFooter>
          </form>
        </Card>
      )}
    </div>
  );
}
