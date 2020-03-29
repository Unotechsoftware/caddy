1. Clone this repository
2. Go to cmd/
3. Run the command go build
4. Verify that the "caddy" executable is built.
5. Run the reverse proxy using the command - < nohup ./caddy run -config sampleConfig/sih.json & >
6. Check logs in /var/log/caddy_default.log


OR

1. Download release
2. It will have binary for Linux x64 and sample configuration


REMEMBER TO CREATE CERTIFICATES - 


1. Also create local CA Certificate and key
2. Sign all certificates with this local CA.
