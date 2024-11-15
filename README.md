# Jerome Otieno's Portfolio

Welcome to my portfolio project! This project showcases my skills and expertise as a software developer, including my proficiency in web development, backend programming, and project management.

## About Me

My name is **Jerome Otieno**, and I am a 19-year-old apprentice at Zone01. I am passionate about creating meaningful and impactful software solutions. My technical expertise includes:
- Programming Languages: **Go, HTML, C#**
- Interests: **Web Development, Backend Development**
- Hobbies: **Basketball, Video Games**

I graduated from:
- **Maseno School** (High School)
- **St Anne's Academy** (Primary School)

## Project Overview

This portfolio is built using modern web technologies and highlights various aspects of my work, including:
- **Contact Form**: A form where visitors can send me a message directly.
- **Skills Section**: A detailed overview of my technical skills and interests.
- **Projects Section**: Showcasing my notable projects, including their descriptions and technologies used.
- **Resume Section**: A downloadable version of my resume.

## Features

- **Responsive Design**: Optimized for both desktop and mobile devices.
- **Contact Form**: Integrated email functionality to receive messages.
- **Dynamic Content**: Sections dynamically loaded using Go and HTML templates.
- **Backend Support**: Powered by the Gin framework in Go.

## Installation

To run this portfolio locally, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/Jerome-afk/my_portfolio.git
    cd portfolio
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Start the server:
    ```bash
    go run main.go
    ```

4. Open your browser and navigate to:
    ```
    http://localhost:4040
    ```

## Technologies Used

- **Backend**: Go (Gin Framework)
- **Frontend**: HTML, CSS, JavaScript
- **Email Integration**: Gomail (SMTP)

## Project Structure

```plaintext
.
├── main.go                # Entry point for the application
├── server/                # Backend server logic
│   └── router.go          # Gin server setup
├── main.html          # Main page
├── error.html         # Error page
├── static/                # Static assets (CSS, JS, images)
├── README.md              # Project documentation
└── go.mod                 # Go module dependencies
