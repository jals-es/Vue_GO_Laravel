<?php

namespace App\Console\Commands;

use App\Jobs\SendEmailJob;
use App\Models\Incidence;
use Illuminate\Console\Command;

class NotifyIncidences extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'send:mail';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Send mail to the admin if there are open incidences.';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return int
     */
    public function handle()
    {
        $total = Incidence::where('status',0)->count();
        if ($total != 0) {
            SendEmailJob::dispatch(env('ADMINMAIL'), 'There are '. $total .' open incidences in the app!');
        }
        return Command::SUCCESS;
    }
}
