#!/bin/bash - 

export FB_LOGIN_URL="https://www.facebook.com/$FB_SDK_VERSION/dialog/oauth?client_id=$FB_CLIENT_ID&redirect_uri=http://$SERVER_ADDR$FB_REDIRECT_URI"

echo -e "Login your Facebook\n$FB_LOGIN_URL\n\n"
echo "Enter authorization code: "
read FB_AUTH_CODE

export FB_GRAPH_URL="https://graph.facebook.com/$FB_SDK_VERSION/oauth/access_token?client_id=$FB_CLIENT_ID&client_secret=$FB_CLIENT_SECRET&redirect_uri=http://$SERVER_ADDR$FB_REDIRECT_URI&code=$FB_AUTH_CODE"

curl $FB_GRAPH_URL

# Debug access token
echo "\nEnter access token to debug: "
read FB_ACCESS_TOKEN

export FB_DEBUG_URL="https://graph.facebook.com/debug_token?input_token=$FB_ACCESS_TOKEN&access_token=$FB_CLIENT_ID|$FB_CLIENT_SECRET\n"

curl $FB_DEBUG_URL
