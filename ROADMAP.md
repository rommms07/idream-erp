# ERP ROADMAP
=============

Before we start implementing new features, we must first finish the designing our login mechanism
for the ERP. For now our main approach are the following:

1. Using Facebook Login to authenticate our user to our ERP.
2. We create a JWT token for users who are authenticated using the manual method.

## Login flow

Our way of authenticating user is quite unique to other platforms. usually we would not require a 
password for identification, there would also be no way of trying to reset a user's password -
Facebook already done that for us!

1. Client wants to sign up or sign in into our app.. Facebook Login will do the job!
2. Is the client already in the system? If it is use authenticate the user directly,
   if it is not ask the user for information about them.

	 - Any valid ID that would verify them as an authenticate user.

   Asking an information about the user from the beginning is mandatory and it is not skipable.
	 This is the case, so that no spam action is performed within the system.

3. (skip when already signed up) After receiving all of the information from the user
   an admin must verify its provided information - here the admin could perform the
	 following actions:

	 - Discard and delete a user request for sign up.
   - Permit a user with a valid information attach with it.
	 - Put a user on hold while no one can verify their identity.

4. There are two ways a user can be signed in the ERP, and it could be one of the following:
   
	 - Facebook Login
	 - Using a generated authentication code by the system.

5. Once they are now authenticated, a token is generated for them - these tokens are then
   used by the system to verify the user.

Note: Facebook Login is mandatory! No one can skip this process during the first phase of sign up.

===================================================================================================

# Graphical representation of what the development roadmap could look like.

<pre>
	+---------------+
  |               |  1. Setting and implementing the app's configuration system. (ok)
	|  Preliminary  |  2. Adding an authentication mechanism.
	|               |  3. Implementing a simple command-line interface that would interact with the system.
	+---------------+
	       ||
				 ||
				 ||
	+---------------+
	|  Information  |  4. Design the gRPC and HTTP middleware mechanism. This is important for further implementation.
	|   Gathering   |  5. 
	+---------------+

</pre>
