API Calculator

Welcome to the API Calculator! This project is a simple yet powerful API built with Go that can handle basic arithmetic operations. Whether you're building a calculator app, integrating with other services, or just experimenting with Go, this API will serve as a great foundation.
🚀 Features

    Basic Arithmetic Operations:

        Addition

        Subtraction

        Multiplication

        Division

        Modulus

    API Endpoints for easy interaction:

        /add : Adds two numbers.

        /subtract : Subtracts the second number from the first.

        /multiply : Multiplies two numbers.

        /divide : Divides the first number by the second (with error handling).

        /modulus : Finds the remainder when dividing two numbers.

🛠️ Technologies Used

    Go (Golang) - A statically typed language designed for simplicity and performance.

🔧 Installation

    Clone this repository:

git clone https://github.com/Loonguinho/GoLangPortifolio.git

Navigate to the project directory:

cd GoLangPortifolio/api-calculator

Install the necessary dependencies:

go mod tidy

Run the server:

    go run main.go

    The API will be running at http://localhost:8080.

📡 API Usage
Endpoints

    Add

        URL: /add

        Method: GET

        Query Params: a (number), b (number)

        Example Request:
            curl --request POST \
              --url http://localhost:8080/add \
              --header 'content-type: application/json' \
              --data '{
              "num1": 5,
              "num2": 3
            }'
        


Response:

    {
      "result": 8
    }

Subtract

    URL: /subtract

    Method: GET

    Query Params: a (number), b (number)

    Example Request:
        curl --request POST \
          --url http://localhost:8080/subtract \
          --header 'content-type: application/json' \
          --data '{
          "num1": 5,
          "num2": 3
        }'


Response:

    {
      "result": 2
    }

Multiply

    URL: /multiply

    Method: GET

    Query Params: a (number), b (number)

    Example Request:
        curl --request POST \
          --url http://localhost:8080/multiply \
          --header 'content-type: application/json' \
          --data '{
          "num1": 5,
          "num2": 3
        }'


Response:

    {
      "result": 15
    }

Divide

    URL: /divide

    Method: GET

    Query Params: a (number), b (number)

    Example Request:
        curl --request POST \
          --url http://localhost:8080/divide \
          --header 'content-type: application/json' \
          --data '{
          "num1": 5,
          "num2": 0
        }'


Response:

    {
      "result": 2
    }

🧑‍💻 Development

    Run the server locally:

        Follow the Installation instructions above to get started.

    Test the API:

        Use any tool like curl, Postman, or directly in the browser to hit the endpoints.

    Contributing:

        Contributions are welcome! If you have suggestions, bug fixes, or improvements, feel free to submit a pull request.

📜 License

This project is open source and available under the MIT License.
