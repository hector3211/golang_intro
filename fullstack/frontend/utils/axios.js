import Axios from "axios";

//axois instance, like a base url for all requests
const instance = Axios.create({
  baseURL: `localhost:8080`,
});

export default instance;
