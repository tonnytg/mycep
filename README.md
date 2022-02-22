# mycep

My Cep é uma API de consulta statica de CEPs.

Os número devem ser digitados sem o traço. Ex: `12345678`

A chamada deverá ser utilizando `/cep` passando o id do CEP como parametro.

        /cep?id=<cep>


Another option is sent `POST` to /cep

    {
        "cep": "12345678"
    }

Exist a example in `request.http` file