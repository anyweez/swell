# Swell: simple configuration management
![Build status](https://travis-ci.org/anyweez/swell.svg?branch=master)

Swell is a simple and less powerful version of popular configuration management
tools like Puppet, Chef, Ansible, and Salt. The goal of Swell is to make it really
easy to ensure that machines meet certain criteria (such as having installed and
up-to-date software running on them, latest versions of configuration files, etc).

When possible, Swell uses prexisting solutions to simply it's own job. In fact, all
it's really doing on it's own is calling a bunch of scripts that are grouped by 
administrators into "profiles", a concept familiar to the other systems above. In
Swell, a profile is simply a directory.

Swell likely won't be adequate for complex setups, but should be far simpler to deploy
on simpler setups than existing tools.

## Quick start
Swell executed scripts that are arranged into profiles (directories), so the first thing you
need to do is to create a directory to house our profiles.

    mkdir profiles
    cd profiles

Edit your `ProfilePath` in `swell.conf` to point to the new directory.

Then create one subdirectory for each profile. Good profiles will capture "packages" of
configurations for end systems; for example `webserver` and `mongodb` are potentially
useful profiles.

    mkdir webserver
    mkdir mongodb
    
Within each directory, add some scripts that should be run for each profile. Note that scripts
should be [idempotent](https://en.wikipedia.org/wiki/Idempotence) since Swell does not try to
get fancy around tracking client state. This means you may want to condition your scripts on 
certain outcomes (for example, don't restart Apache unless something happened that'd make you
need to restart Apache, like downloading a new configuration).

After you've added scripts, there are two steps remaining: make the `profiles` directory accessible
to each client machine (network mounts, FTP, wget, etc) and configure some sort of trigger to
invoke `swell`. I usually use cron, but many different utilities are available for triggering on
events other than time if desired. Run Swell with the profiles that the machine should apply and
Swell will ensure that all scripts are run.

    swell webserver mongodb
    
And you're done!