<?php

namespace App\Mail;

use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Mail\Mailable;
use Illuminate\Queue\SerializesModels;

class AdminNotification extends Mailable
{
    use Queueable, SerializesModels;
    protected $content;

    /**
     * Create a new message instance.
     *
     * @return void
     */
    public function __construct($content = "New Incidence registered in the app!")
    {
        $this->content = $content;
        //
    }

    /**
     * Build the message.
     *
     * @return $this
     */
    public function build()
    {
        return $this->from('notify@appbar.com', 'APPBAR ADMIN')
        ->subject('Open incidence')
        ->html($this->content);
    }
}
