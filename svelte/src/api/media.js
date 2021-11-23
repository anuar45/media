// pokemon.js
// Implementations for all the calls for the pokemon endpoints.
import Api from "../services/Api";

// Method to get a list of all items in a path
export const getMediaItems = async (path) => {
    try {
      const response = await Api.get("/items?path=" + path);
      return response;
    } catch (error) {
      console.log(error)
    }
};

// Get a pokemon details by name
// export const getPokemonByName = async(name) => {
//     try {
//       const response = await Api.get(`/pokemon/${name}`);
//       return response;
//     } catch (error) {
//       console.error(error);
//     }
// };
