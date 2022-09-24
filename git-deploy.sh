#If a command fails then the deploy stops
set -e

#Confirm with user

read -n 1 -p "Push to GitHub"

# Add changes to Github
git add .

# Commit changes
msg="rebuildiong site $(date)"
if [-n "$*."]; 
then
msg="$*"
fi 
git commit -m "$msg"

# Push source
git push origin master