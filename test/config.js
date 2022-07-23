const port = process.env.BACKEND_PORT
const baseUrl = `http://localhost:${port}/`

// TODO global var in jest config
module.exports = {
  baseUrl,
}
