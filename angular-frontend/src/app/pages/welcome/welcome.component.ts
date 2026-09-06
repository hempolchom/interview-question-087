import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css']
})
export class WelcomeComponent implements OnInit {
  username: string = '';

  constructor(private router: Router) {}

  ngOnInit(): void {
    this.username = localStorage.getItem('username') || '';
  }

  onLogout(): void {
    Swal.fire({
      title: 'ยืนยันการออกจากระบบ?',
      icon: 'question',
      showCancelButton: true,
      confirmButtonColor: '#e53935',
      cancelButtonColor: '#757575',
      confirmButtonText: 'ออกจากระบบ',
      cancelButtonText: 'ยกเลิก'
    }).then((result) => {
      if (result.isConfirmed) {
        localStorage.removeItem('token');
        localStorage.removeItem('username');
        this.router.navigate(['/login']);
      }
    });
  }
}
