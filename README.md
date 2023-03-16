# Beam

Beam is a payment solution that allows businesses to easily create and manage checkout links for content or services using the Lightning Network. The paywall will be a checkout link that customers can share directly. Customers will be prompted to pay the required amount in Bitcoin through the Lightning Network when they click the link.

Users can be individuals, businesses, shops, or online vendors. The app will enable users to generate web urls, and track payments made to that url.

Functionalities

- Generate payment links
- View payment links
- Track payments made on the platform

Tools:
Programming language: Golang. Backend server: Echo framework. Database: PostgreSQL Deployment: Docker, Docker Compose.

Dependency software:

- Polar
- [lnrpc](github.com/lncm/lnd-rpc/v0.10.0/lnrpc)

Dependency libraries:
Lnrpc for interacting with lnd. sqlc for database operations.

Backend:
RESTful API using Echo framework. Modular file structure. Endpoints:

The payment link is generated and added to the database, this payment link can be sent to multiple people for payment for things.
