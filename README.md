# Journalful – Academic Journal Logging and Tracking Tool

Journalful is a web-based application designed to help academics, researchers, and students keep track of their reading of academic journals. It provides a simple interface to save articles, track reading progress, and manage a personal library of academic literature.

## Features

*   **Article Management:** Save articles with their DOI, title, authors, and other relevant information.
*   **Library Management:** Organize your saved articles into a personal library.
*   **Reading Status:** Track the status of each article (e.g., To Read, Reading, Read).
*   **User Profiles:** Manage your user profile and preferences.
*   **gRPC API:** A modern, high-performance API for all application functionalities.
*   **Web Interface:** A user-friendly web interface built with Nuxt.js.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

*   [Go](https://golang.org/) (version 1.20 or later)
*   [Node.js](https://nodejs.org/) (version 18 or later)
*   [Docker](https://www.docker.com/)
*   [Docker Compose](https://docs.docker.com/compose/)
*   [Make](https://www.gnu.org/software/make/)

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/chiquitav2/journalful.git
    cd journalful
    ```

2.  **Run the application:**

    ```bash
    make up
    ```

    This command will build the Go application, the Nuxt.js UI, and start the necessary services using Docker Compose.

3.  **Access the application:**

    The web interface will be available at [http://localhost:3000](http://localhost:3000).

## Usage

Once the application is running, you can create an account, start searching for articles, and add them to your library. You can then update the reading status of each article as you progress.

## API Overview

The application exposes a gRPC API for all its functionalities. The API is defined in the `.proto` files located in the `api` directory. For more details on the API, please refer to the `API.md` file.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or find any bugs.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.