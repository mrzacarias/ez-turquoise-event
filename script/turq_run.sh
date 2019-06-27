#!/bin/sh

set -e

cd ~/everquote/turquoise && npm run trigger && cd ~/everquote/ez-turquoise-event
