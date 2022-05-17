// darood crore

const sqlite3 = require('sqlite3').verbose();
//const db = new sqlite3.Database('darood2.db');

const express = require('express')
const app = express()
const port = 3010

var counter = 0;
app.set('view engine', 'ejs');

app.get('/', (req, res)=> {
    res.send("crore darood");
});

app.use('/static', express.static('/static'));

app.get('/users/:name', (req, res) => {
  var name = req.params.name || 'ikb'
  // res.send('Hello name:' + name)
  res.render("user",{name:name})
})

app.put('/users/:name/counters/incr', (req, res) => {
  var name = req.params.name || 'ikb'
  // res.send('Hello name:' + name)
  var stepStr = req.query.step || '1'
  var step = parseInt(stepStr) || 1;
  counter = counter + step;
  res.json({counter: counter})
})

app.get('/users/:name/counters', (req, res) => {
  var name = req.params.name || 'ikb'
  // res.send('Hello name:' + name)

    const db = new sqlite3.Database('darood2.db');
    const sql= 'select * from counters where username=?';
    db.get(sql, name, (err, result) => {
	if(err){
	    res.json({err:err});
	}else{
	    res.json({result: result});
	}
    });
    db.close();
  res.json({counter: counter});
})

app.post('/users/:name/counters',(req,res) => {
    console.log("inside create counter");
    //todo create counter table in darood2.db sqlite3
    // create table counters(id integer primary key autoincrement, username varchar(512), value integer)
    const db = new sqlite3.Database('daroo2.db');
    const username = req.params.name;
    const stmt = db.prepare("insert into counters(username) values(?)");
    stmt.run(username);
    stmt.finalize();
    db.close();
    res.json({status: "counter created"});
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})
