package templates

templ Layout() {
  <!DOCTYPE html>
  <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Go Website</title>
        <script src="https://cdn.tailwindcss.com"></script>
        <script src="https://unpkg.com/htmx.org"></script>
        <script src="https://cdn.jsdelivr.net/npm/alpinejs" defer></script>
    </head>
    <body class="bg-gray-100 text-gray-800">
      @Navbar()
      <main class="container mx-auto mt-4">
        {children...}
      </main>
      @Footer()
    </body>
  </html>
}
