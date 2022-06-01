use crossbeam_channel::{bounded, unbounded, Receiver, Sender};
use crossterm::event::{Event, KeyCode, KeyEvent, KeyModifiers};

use lazy_static::lazy_static;

use std::sync::atomic::AtomicBool;
use std::sync::{Arc, Mutex};
use std::time::Duration;


lazy_static! {
    pub static ref NOTIFY: (Sender<HGEvent>, Receiver<HGEvent>) = bounded(1024);
    pub static ref GG_COMBINE: AtomicBool = AtomicBool::new(false);
}

pub struct App {
    pub mode: AppMode,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum AppMode {
    /// 搜索模式
    Search,

    /// 浏览模式
    View,

    /// 弹窗提示
    Popup,

    /// 项目明细
    Detail,
}

#[derive(Debug, Clone)]
pub enum HGEvent {
    UserEvent(KeyEvent),

    NotifyEvent(Notify),
}

#[derive(Debug, Clone, PartialEq)]
pub enum Notify {
    /// 重绘界面
    Redraw,

    /// 退出应用
    Quit,

    /// 弹出窗口展示消息
    Message(Message),

    /// tick
    Tick,
}


#[derive(Debug, Clone, PartialEq)]
pub enum Message {
    Error(String),

    Warn(String),

    Tips(String),
}

// default接口，设置默认值

impl Default for Message {
    fn default() -> Self {
        Message::Error(String::default())
    }
}


pub fn handle_key_event(event_app: Arc<Mutex<App>>) {
    let (sender, receiver) = unbounded();
    sender
        .send(HGEvent::UserEvent(KeyEvent {
            code: KeyCode::Char('#'),
            modifiers: KeyModifiers::NONE,
        }))
        .unwrap();

    sender
        .send(HGEvent::UserEvent(KeyEvent {
            code: KeyCode::Enter,
            modifiers: KeyModifiers::NONE,
        }))
        .unwrap();


    // 将读写操作都加入到额外的线程中
    std::thread::spawn(move || loop {
        if let Ok(Event::Key(event)) = crossterm::event::read() {
            sender.send(HGEvent::UserEvent(event)).unwrap();
        }
    });
    std::thread::spawn(move || loop {
        if let Ok(HGEvent::UserEvent(key_event)) = receiver.recv() {
            let mut app = event_app.lock().unwrap();
            match (key_event.modifiers, key_event.code) {
                (KeyModifiers::CONTROL, KeyCode::Char('c')) => {
                    quit();
                    break;
                }
                (key_modifier, key_code) => match app.mode {
                    AppMode::Search => {
                        println!("search")
                    }
                    AppMode::View => {
                        println!("view")
                    }
                    AppMode::Popup => {
                        println!("popup")
                    }
                    AppMode::Detail => {
                        println!("detail")
                    }
                },
            }
        }
    });
}

pub fn redraw() {
    NOTIFY.0.send(HGEvent::NotifyEvent(Notify::Redraw)).unwrap();
}

pub fn quit() {
    println!("quit")
}