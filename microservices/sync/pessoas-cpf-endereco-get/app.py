from flask import Flask, Response, jsonify

app = Flask(__name__)

result = {
        'endereco': 'Rua Joao Albuquerque',
        'numero': 123,
        'complemento': 'apto 45', 
        'bairro': 'Ipiranga',
        'cidade': 'São Paulo',
        'estado': 'SP'
    }

@app.route('/enderecos/<cpf>', methods=['GET'])
def get_endereco(cpf):
    if int(cpf) == 123:
        return jsonify({ 'data': result })
    else:
        return Response('', status=404, mimetype='application/json')

if __name__ == '__main__':
    app.run(debug=True, port=3000)