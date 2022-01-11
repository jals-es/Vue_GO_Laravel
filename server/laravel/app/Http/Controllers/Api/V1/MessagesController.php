<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Models\Message;
use App\Events\NewMessage;
use App\Traits\UserTrait;
use Pusher\Pusher;

class MessagesController extends Controller
{
    use UserTrait;
    public function fetch()
    {
        return response()->json(['status' => 'success', 'data' => Message::all()], 201);
    }
    public function store()
    {


        Message::create([
            'content' => request()->content,
            'author' => self::getId(),
            'author' => "antoni",
        ]);
        $options = array(
            'cluster' => env('PUSHER_APP_CLUSTER'),
            'encrypted' => true
        );
        $pusher = new Pusher(
            env('PUSHER_APP_KEY'),
            env('PUSHER_APP_SECRET'),
            env('PUSHER_APP_ID'),
            $options
        );
        $data['message'] = request()->content;
        $data['author'] = self::getId();
        $pusher->trigger('messages', 'App\\Events\\NewMessage', $data);
        return ['status' => 'Message Sent!'];
    }
}
