# time_wrapper - Execute or don't execute commands based on current time

### Motivation

One way to not be woken up at night is to downgrade your checks
from critical to warning, another way is to only alert during office
hours. This doesn't always make sense, but for example if you prefer
to run a cronjob during the night, but if it fails you only want to be
alerted in the morning - either configure your monitoring, or wrap the
check with this tool.

### Usage

  * Example 1
    * `time_wrapper --days 1,2,3,4,5 --from 8 --to 18 -- date`
    * Execute `date`, but only Mon-Fri from 08:00:00 to 17:59:59 UTC
  * Example 2
    * `time_wrapper --days 0,6 --from 20 --to 24 -- date +%s`
    * Execute `date +%s`, but only Sat+Sun from 20:00:00 to 23:59:59 UTC
  * Help
    * `time_wrapper -h`

### License

Copyright (C) 2016-2017 wywy GmbH

Permission to use, copy, modify, and/or distribute this software for any purpose with or without fee is hereby granted.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
