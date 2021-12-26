<?php

namespace App\Jobs;

use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldBeUnique;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Foundation\Bus\Dispatchable;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Queue\SerializesModels;
use App\Mail\AdminNotification;
use Illuminate\Support\Facades\Mail;
class SendEmailJob implements ShouldQueue
{
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;
    protected $send_mail;
    protected $content;

    /**
     * Create a new job instance.
     *
     * @return void
     */
    public function __construct($send_mail, $content = "New Incidence registered in the app!")
    {
        $this->content = $content;
        $this->send_mail = $send_mail;
    }

    /**
     * Execute the job.
     *
     * @return void
     */
    public function handle()
    {
        $this->onQueue('processing');
        $email = new AdminNotification($this->content);
        Mail::to('antonitormo@gmail.com')->send($email);
        $this->release();

    }
}
