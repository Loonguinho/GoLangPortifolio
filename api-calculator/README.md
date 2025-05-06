# HTMX Calculator Exercise (Go Backend)

## Objective

Build a calculator application where:

- Frontend uses **HTMX** for all interactions (zero custom JavaScript)
- Backend in **Go** performs all calculations
- Communication via HTML-over-HTTP

## Backend Requirements (`main.go`)

### Endpoints

1. `GET /`
   - Serves static HTML page
2. `POST /calculate`
   - Accepts `application/x-www-form-urlencoded`
   - Processes: `+`, `-`, `*`, `/`
   - Returns HTML fragments
3. `GET /history` (Bonus)
   - Returns calculation history

### Must Handle

- Input validation (numeric, division by zero)
- Proper HTTP status codes
- Clean error messages in HTML format
- Memory-only storage (no DB required)

## Frontend Requirements (`index.html`)

### HTMX Features

```html
<form hx-post="/calculate" hx-target="#result" hx-swap="innerHTML">
  <!-- Input fields -->
</form>
```
