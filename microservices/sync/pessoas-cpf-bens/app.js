const http = require('http');
const express = require('express');
const Router = express.Router();

const app = express();

Router['get']('/pessoas/:cpf/bens', (req, res) => {
    if (req.params.cpf === '123') {
        const result = [
            {
                tipo: 'casa',
                valor: 300.000
            },
            {
                tipo: 'veiculo',
                valor: 70.000
            }
        ];
        res.status(200).send({ data: result });
        return;
    }
    res.status(204).send();
});

app.use(Router);
app.use((req, res) => {
    res.status(404).send();
});

server = http.createServer(app);
server.listen(3000, () => {
    console.log(`Servidor iniciado em: localhost:3000`);
})