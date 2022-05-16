// darood crore


const express = require('express')
const app = express()
const port = 3010


app.set('view engine', 'ejs');

app.use('/static', express.static('/static'));

app.get('/users/:name', (req, res) => {
  var name = req.params.name || 'ikb'
  // res.send('Hello name:' + name)
  res.render("user",{name:name})
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})
