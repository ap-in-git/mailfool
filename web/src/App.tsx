import {createBrowserRouter, RouterProvider,} from "react-router-dom";
import Home from "./pages/Home";
import SignIn from "./pages/SignIn";

const router = createBrowserRouter([
    {
        path:"/",
        element: <Home/>
    },
    {
        path:"/login",
        element:<SignIn/>
    }

])
function App() {
  return (
      <RouterProvider router={router}/>
  )
}

export default App
