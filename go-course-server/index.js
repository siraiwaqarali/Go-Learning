const express = require('express')
const app = express()
const port = 3000

app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.get('/', (req, res) => {
    res.status(200).send("Welcome to Go Course Learning Server")
})

app.get('/get', (req, res) => {
    res.status(200).json({ message: "Hello from CodeWithWaqar" })
})

app.post('/post', (req, res) => {
    let reqBody = req.body
    res.status(200).send(reqBody)
})

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body))
})

app.listen(port, () => {
    console.log(`Server is running on port ${port}`)
})