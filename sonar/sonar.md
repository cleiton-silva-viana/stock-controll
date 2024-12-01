docker run \
    --rm \
    -v "C:\Users\Inara\Documents\stock-controll\internal:/usr/src" \
    --network="host" \
    -e SONAR_HOST_URL="http://localhost:9000" \
    -e SONAR_SCANNER_OPTS="-Dsonar.projectKey=ecommerce" \
    -e SONAR_TOKEN="sqp_591c72dffcbf0d24cc244bf8c3510d99d13546c2" \
    sonarsource/sonar-scanner-cli