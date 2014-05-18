# Nudger 

Nudger is a very simple webservice that allows triggering an action ("nudge") on a limited set of items via HTTP.

The action can be any command-line, pre-set in the config.
The permissible set of items is given by the entries of a given directory.


# Use as a web hook handler

An example application is the automatic refreshing of cvs code via a web hook.
Example: someone pushed a commit to master on a gitlab server, this triggers a web hook directed at this application, which uses a "git pull" command to update the code to the latest version.

# API's

## GET /

Gives you a brief help text

## GET /items

Gives the list of permissible items

## GET /items/:item

Nudges (executes the action), passing :item as argument.

## POST /items/:item

Same as GET /items/:item